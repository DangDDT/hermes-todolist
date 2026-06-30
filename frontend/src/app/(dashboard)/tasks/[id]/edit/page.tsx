'use client';

import { useParams } from 'next/navigation';
import Link from 'next/link';
import { ArrowLeft, AlertCircle, RefreshCw } from 'lucide-react';
import { useTask, useUpdateTask } from '@/features/tasks/hooks';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Skeleton } from '@/components/ui/skeleton';

export default function EditTaskPage() {
  const params = useParams();
  const id = params.id as string;
  const { data: taskData, isLoading, isError, error, refetch } = useTask(id);
  const updateTask = useUpdateTask(id);

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
            Make changes to your task — full form will be implemented with react-hook-form + zod
          </CardDescription>
        </CardHeader>
        <CardContent className="flex flex-col items-center justify-center py-12">
          <div className="rounded-full bg-muted p-4 mb-4">
            <ArrowLeft className="h-8 w-8 text-muted-foreground" />
          </div>
          <h3 className="text-lg font-semibold mb-2">Edit Form Coming Soon</h3>
          <p className="text-muted-foreground text-center max-w-md mb-4">
            The full edit form with react-hook-form + zod validation will be implemented
            in the next iteration.
          </p>
          <Button variant="outline" render={<Link href={`/tasks/${id}`} />}>
            Back to task
          </Button>
        </CardContent>
      </Card>
    </div>
  );
}
