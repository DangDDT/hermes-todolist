'use client';

import { useState, type FormEvent } from 'react';
import { useParams } from 'next/navigation';
import Link from 'next/link';
import { ArrowLeft, Pencil, Trash2, AlertCircle, RefreshCw, MessageSquare, Send } from 'lucide-react';
import { useTask, useDeleteTask } from '@/features/tasks/hooks';
import { useTaskComments, useCreateTaskComment } from '@/features/task-comments/hooks';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Skeleton } from '@/components/ui/skeleton';
import { Separator } from '@/components/ui/separator';
import { Textarea } from '@/components/ui/textarea';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import { cn, formatDateTime } from '@/lib/utils';

const priorityColors: Record<string, string> = {
  low: 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200',
  medium: 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-200',
  high: 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200',
};

const statusColors: Record<string, string> = {
  todo: 'bg-gray-100 text-gray-800 dark:bg-gray-800 dark:text-gray-200',
  in_progress: 'bg-purple-100 text-purple-800 dark:bg-purple-900 dark:text-purple-200',
  done: 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200',
};

export default function TaskDetailPage() {
  const params = useParams();
  const id = params.id as string;
  const { data, isLoading, isError, error, refetch } = useTask(id);
  const { data: commentsData, isLoading: commentsLoading, isError: commentsIsError, error: commentsError, refetch: refetchComments } = useTaskComments(id);
  const createComment = useCreateTaskComment(id);
  const deleteTask = useDeleteTask();
  const [showDeleteDialog, setShowDeleteDialog] = useState(false);
  const [commentBody, setCommentBody] = useState('');

  if (isLoading) {
    return (
      <div className="max-w-2xl mx-auto space-y-6">
        <Skeleton className="h-9 w-24" />
        <Card>
          <CardHeader>
            <Skeleton className="h-7 w-3/4" />
            <Skeleton className="h-4 w-1/2 mt-2" />
          </CardHeader>
          <CardContent className="space-y-4">
            <Skeleton className="h-4 w-full" />
            <Skeleton className="h-4 w-5/6" />
            <Skeleton className="h-4 w-4/6" />
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

  if (!data?.task) {
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

  const task = data.task;

  function handleCommentSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const body = commentBody.trim();
    if (!body) return;
    createComment.mutate(
      { body },
      {
        onSuccess: () => {
          setCommentBody('');
        },
      }
    );
  }

  function handleDelete() {
    deleteTask.mutate(id);
    setShowDeleteDialog(false);
  }

  return (
    <div className="max-w-2xl mx-auto space-y-6">
      {/* Navigation */}
      <div className="flex items-center justify-between">
        <Button variant="ghost" size="sm" render={<Link href="/tasks" />}>
          <ArrowLeft className="mr-2 h-4 w-4" />
          Back to tasks
        </Button>
        <div className="flex gap-2">
          <Button variant="outline" size="sm" render={<Link href={`/tasks/${id}/edit`} />}>
            <Pencil className="mr-2 h-4 w-4" />
            Edit
          </Button>
          <Button
            variant="destructive"
            size="sm"
            onClick={() => setShowDeleteDialog(true)}
          >
            <Trash2 className="mr-2 h-4 w-4" />
            Delete
          </Button>
        </div>
      </div>

      {/* Task card */}
      <Card>
        <CardHeader>
          <div className="flex items-start justify-between gap-4">
            <div className="flex-1">
              <CardTitle className={cn('text-2xl', task.status === 'done' && 'line-through')}>
                {task.title}
              </CardTitle>
              <div className="flex items-center gap-2 mt-2">
                <Badge className={cn(statusColors[task.status])}>
                  {task.status.replace('_', ' ')}
                </Badge>
                <Badge className={cn(priorityColors[task.priority])}>
                  {task.priority}
                </Badge>
              </div>
            </div>
          </div>
        </CardHeader>
        <CardContent className="space-y-6">
          {task.description ? (
            <div>
              <h3 className="text-sm font-medium text-muted-foreground mb-2">Description</h3>
              <p className="text-sm whitespace-pre-wrap">{task.description}</p>
            </div>
          ) : (
            <p className="text-sm text-muted-foreground italic">No description provided</p>
          )}

          <Separator />

          <div className="grid gap-4 sm:grid-cols-2">
            {task.dueDate && (
              <div>
                <h3 className="text-sm font-medium text-muted-foreground mb-1">Due Date</h3>
                <p className="text-sm">{formatDateTime(task.dueDate)}</p>
              </div>
            )}
            <div>
              <h3 className="text-sm font-medium text-muted-foreground mb-1">Created</h3>
              <p className="text-sm">{formatDateTime(task.createdAt)}</p>
            </div>
            <div>
              <h3 className="text-sm font-medium text-muted-foreground mb-1">Last Updated</h3>
              <p className="text-sm">{formatDateTime(task.updatedAt)}</p>
            </div>
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2 text-lg">
            <MessageSquare className="h-5 w-5" />
            Comments
          </CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <form className="space-y-3" onSubmit={handleCommentSubmit}>
            <Textarea
              value={commentBody}
              onChange={(event) => setCommentBody(event.target.value)}
              placeholder="Write a comment to update the team…"
              rows={4}
            />
            <div className="flex items-center justify-between gap-3">
              <p className="text-xs text-muted-foreground">
                Use comments for review notes, clarification, or handoff context.
              </p>
              <Button type="submit" disabled={createComment.isPending || !commentBody.trim()}>
                <Send className="mr-2 h-4 w-4" />
                {createComment.isPending ? 'Posting...' : 'Post comment'}
              </Button>
            </div>
          </form>

          <Separator />

          {commentsLoading ? (
            <div className="space-y-3">
              <Skeleton className="h-16 w-full" />
              <Skeleton className="h-16 w-full" />
            </div>
          ) : commentsIsError ? (
            <div className="rounded-lg border border-destructive/40 bg-destructive/5 p-4 text-sm">
              <p className="font-medium text-destructive">Failed to load comments</p>
              <p className="mt-1 text-muted-foreground">
                {commentsError?.message || 'An unexpected error occurred'}
              </p>
              <Button variant="outline" size="sm" className="mt-3" onClick={() => refetchComments()}>
                <RefreshCw className="mr-2 h-4 w-4" />
                Retry comments
              </Button>
            </div>
          ) : (commentsData?.comments?.length ?? 0) === 0 ? (
            <div className="rounded-lg border border-dashed border-border bg-muted/20 p-6 text-sm text-muted-foreground">
              No comments yet. Be the first one to add context for this task.
            </div>
          ) : (
            <div className="space-y-3">
              {commentsData!.comments.map((comment) => (
                <div key={comment.id} className="rounded-lg border bg-muted/20 p-4">
                  <div className="flex items-center justify-between gap-3">
                    <div className="font-medium">{comment.authorName || 'Team member'}</div>
                    <div className="text-xs text-muted-foreground">{formatDateTime(comment.createdAt)}</div>
                  </div>
                  <p className="mt-2 whitespace-pre-wrap text-sm leading-6">{comment.body}</p>
                </div>
              ))}
            </div>
          )}
        </CardContent>
      </Card>

      {/* Delete confirmation dialog */}
      <Dialog open={showDeleteDialog} onOpenChange={setShowDeleteDialog}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Delete Task</DialogTitle>
            <DialogDescription>
              Are you sure you want to delete &ldquo;{task.title}&rdquo;? This action cannot be undone.
            </DialogDescription>
          </DialogHeader>
          <DialogFooter showCloseButton>
            <Button
              variant="destructive"
              onClick={handleDelete}
              disabled={deleteTask.isPending}
            >
              {deleteTask.isPending ? 'Deleting...' : 'Delete'}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
