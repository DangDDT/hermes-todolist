'use client';

import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import Link from 'next/link';
import { ArrowLeft } from 'lucide-react';
import { CreateTaskSchema, type CreateTaskInput } from '@/lib/zod-schemas';
import { useCreateTask } from '@/features/tasks/hooks';
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

export default function NewTaskPage() {
  const createTask = useCreateTask();

  const {
    register,
    handleSubmit,
    setValue,
    watch,
    formState: { errors, isSubmitting },
  } = useForm<CreateTaskInput>({
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    resolver: zodResolver(CreateTaskSchema) as any,
    defaultValues: {
      priority: 'medium',
      status: 'todo',
    },
  });

  const priority = watch('priority');
  const status = watch('status');

  function onSubmit(data: CreateTaskInput) {
    createTask.mutate(data);
  }

  return (
    <div className="max-w-2xl mx-auto space-y-6">
      {/* Page header */}
      <div className="flex items-center gap-4">
        <Button variant="ghost" size="icon" render={<Link href="/tasks" />}>
          <ArrowLeft className="h-5 w-5" />
          <span className="sr-only">Back to tasks</span>
        </Button>
        <div>
          <h1 className="text-2xl font-bold tracking-tight">Create Task</h1>
          <p className="text-muted-foreground mt-1">
            Add a new task to your list
          </p>
        </div>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Task Details</CardTitle>
          <CardDescription>
            Fill in the details for your new task
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
                <Select
                  value={priority}
                  onValueChange={(value) =>
                    setValue('priority', value as CreateTaskInput['priority'], {
                      shouldValidate: true,
                    })
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
                {errors.priority && (
                  <p className="text-sm text-destructive">{errors.priority.message}</p>
                )}
              </div>

              <div className="space-y-2">
                <label htmlFor="status" className="text-sm font-medium">
                  Status
                </label>
                <Select
                  value={status}
                  onValueChange={(value) =>
                    setValue('status', value as CreateTaskInput['status'], {
                      shouldValidate: true,
                    })
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
                disabled={isSubmitting || createTask.isPending}
              >
                {createTask.isPending ? 'Creating...' : 'Create Task'}
              </Button>
              <Button type="button" variant="outline" render={<Link href="/tasks" />}>
                Cancel
              </Button>
            </div>

            {createTask.isError && (
              <div className="rounded-md bg-destructive/15 p-3 text-sm text-destructive">
                {createTask.error?.message || 'Failed to create task'}
              </div>
            )}
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
