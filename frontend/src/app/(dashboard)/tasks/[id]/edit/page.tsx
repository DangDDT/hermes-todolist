'use client';

import { useForm, Controller } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useParams } from 'next/navigation';
import Link from 'next/link';
import { ArrowLeft, AlertCircle, RefreshCw } from 'lucide-react';
import { UpdateTaskSchema, type UpdateTaskInput } from '@/lib/zod-schemas';
import { useTask, useUpdateTask } from '@/features/tasks/hooks';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Skeleton } from '@/components/ui/skeleton';

export default function EditTaskPage() {
  const params = useParams();
  const id = params.id as string;
  const { data: taskData, isLoading, isError, error, refetch } = useTask(id);
  const updateTask = useUpdateTask(id);

  const {
    register,
    handleSubmit,
    control,
    formState: { errors, isSubmitting },
  } = useForm<UpdateTaskInput>({
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    resolver: zodResolver(UpdateTaskSchema) as any,
    values: taskData?.task
      ? {
          title: taskData.task.title,
          description: taskData.task.description || '',
          priority: taskData.task.priority,
          status: taskData.task.status,
          dueDate: taskData.task.dueDate || null,
        }
      : undefined,
  });

  function onSubmit(data: UpdateTaskInput) {
    // Clean undefined values for PATCH
    const cleaned: UpdateTaskInput = {};
    if (data.title !== undefined) cleaned.title = data.title;
    if (data.description !== undefined) cleaned.description = data.description || undefined;
    if (data.priority !== undefined) cleaned.priority = data.priority;
    if (data.status !== undefined) cleaned.status = data.status;
    if (data.dueDate !== undefined) cleaned.dueDate = data.dueDate || null;
    updateTask.mutate(cleaned);
  }

  if (isLoading) {
    return (
      <div className="max-w-2xl mx-auto space-y-6">
        <Skeleton className="h-9 w-24" />
        <Card>
          <CardHeader>
            <Skeleton className="h-7 w-3/4" />
            <Skeleton className="h-4 w-1/2 mt-2" />
          </CardHeader>
          <CardContent className="space-y-6">
            <Skeleton className="h-12 w-full" />
            <Skeleton className="h-24 w-full" />
            <div className="grid grid-cols-3 gap-4">
              <Skeleton className="h-12 w-full" />
              <Skeleton className="h-12 w-full" />
              <Skeleton className="h-12 w-full" />
            </div>
          </CardContent>
        </Card>
      </div>
    );
  }

  if (isError) {
    return (
      <div className="max-w-2xl mx-auto space-y-6">
        <Card className="border-destructive">
          <CardContent className="flex flex-col items-center justify-center py-12">
            <AlertCircle className="h-12 w-12 text-destructive mb-4" />
            <h3 className="text-lg font-semibold mb-2">Failed to load task</h3>
            <p className="text-muted-foreground mb-4">
              {error?.message || 'An unexpected error occurred'}
            </p>
            <Button variant="outline" onClick={() => refetch()}>
              <RefreshCw className="mr-2 h-4 w-4" />
              Retry
            </Button>
          </CardContent>
        </Card>
      </div>
    );
  }

  if (!taskData?.task) {
    return (
      <div className="max-w-2xl mx-auto space-y-6">
        <Card>
          <CardContent className="flex flex-col items-center justify-center py-12">
            <h3 className="text-lg font-semibold mb-2">Task not found</h3>
            <p className="text-muted-foreground mb-4">
              This task may have been deleted or doesn&apos;t exist.
            </p>
            <Button render={<Link href="/tasks" />}>
              Back to tasks
            </Button>
          </CardContent>
        </Card>
      </div>
    );
  }

  return (
    <div className="max-w-2xl mx-auto space-y-6">
      {/* Page header */}
      <div className="flex items-center gap-4">
        <Button variant="ghost" size="icon" render={<Link href={`/tasks/${id}`} />}>
          <ArrowLeft className="h-5 w-5" />
          <span className="sr-only">Back to task</span>
        </Button>
        <div>
          <h1 className="text-2xl font-bold tracking-tight">Edit Task</h1>
          <p className="text-muted-foreground mt-1">
            Update your task details
          </p>
        </div>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Task Details</CardTitle>
          <CardDescription>
            Make changes to your task
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
            <div className="space-y-2">
              <label htmlFor="title" className="text-sm font-medium">
                Title <span className="text-destructive">*</span>
              </label>
              <Input
                id="title"
                placeholder="What needs to be done?"
                {...register('title')}
                aria-invalid={!!errors.title}
              />
              {errors.title && (
                <p className="text-sm text-destructive">{errors.title.message}</p>
              )}
            </div>

            <div className="space-y-2">
              <label htmlFor="description" className="text-sm font-medium">
                Description
              </label>
              <Textarea
                id="description"
                placeholder="Add more details about this task..."
                rows={4}
                {...register('description')}
                aria-invalid={!!errors.description}
              />
              {errors.description && (
                <p className="text-sm text-destructive">{errors.description.message}</p>
              )}
            </div>

            <div className="grid gap-4 sm:grid-cols-3">
              <div className="space-y-2">
                <label htmlFor="priority" className="text-sm font-medium">
                  Priority
                </label>
                <Controller
                  name="priority"
                  control={control}
                  render={({ field }) => (
                    <Select
                      value={field.value || 'medium'}
                      onValueChange={(value) =>
                        field.onChange(value as UpdateTaskInput['priority'])
                      }
                    >
                      <SelectTrigger id="priority">
                        <SelectValue placeholder="Select priority" />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="low">Low</SelectItem>
                        <SelectItem value="medium">Medium</SelectItem>
                        <SelectItem value="high">High</SelectItem>
                      </SelectContent>
                    </Select>
                  )}
                />
                {errors.priority && (
                  <p className="text-sm text-destructive">{errors.priority.message}</p>
                )}
              </div>

              <div className="space-y-2">
                <label htmlFor="status" className="text-sm font-medium">
                  Status
                </label>
                <Controller
                  name="status"
                  control={control}
                  render={({ field }) => (
                    <Select
                      value={field.value || 'todo'}
                      onValueChange={(value) =>
                        field.onChange(value as UpdateTaskInput['status'])
                      }
                    >
                      <SelectTrigger id="status">
                        <SelectValue placeholder="Select status" />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="todo">To Do</SelectItem>
                        <SelectItem value="in_progress">In Progress</SelectItem>
                        <SelectItem value="done">Done</SelectItem>
                      </SelectContent>
                    </Select>
                  )}
                />
                {errors.status && (
                  <p className="text-sm text-destructive">{errors.status.message}</p>
                )}
              </div>

              <div className="space-y-2">
                <label htmlFor="dueDate" className="text-sm font-medium">
                  Due Date
                </label>
                <Input
                  id="dueDate"
                  type="date"
                  {...register('dueDate')}
                  aria-invalid={!!errors.dueDate}
                />
                {errors.dueDate && (
                  <p className="text-sm text-destructive">{errors.dueDate.message}</p>
                )}
              </div>
            </div>

            <div className="flex gap-4 pt-4">
              <Button
                type="submit"
                disabled={isSubmitting || updateTask.isPending}
              >
                {updateTask.isPending ? 'Saving...' : 'Save Changes'}
              </Button>
              <Button type="button" variant="outline" render={<Link href={`/tasks/${id}`} />}>
                Cancel
              </Button>
            </div>

            {updateTask.isError && (
              <div className="rounded-md bg-destructive/15 p-3 text-sm text-destructive">
                {updateTask.error?.message || 'Failed to update task'}
              </div>
            )}
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
