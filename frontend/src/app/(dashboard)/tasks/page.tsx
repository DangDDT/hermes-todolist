'use client';

import { useState } from 'react';
import Link from 'next/link';
import { Plus, Search, AlertCircle, RefreshCw } from 'lucide-react';
import { useTasks } from '@/features/tasks/hooks';
import { useDeleteTask } from '@/features/tasks/hooks';
import type { Task } from '@/features/tasks/types';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import { Skeleton } from '@/components/ui/skeleton';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import { cn, formatDate } from '@/lib/utils';

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

export default function TasksPage() {
  const [search, setSearch] = useState('');
  const [statusFilter, setStatusFilter] = useState<string | undefined>();
  const { data, isLoading, isError, error, refetch } = useTasks({
    search: search || undefined,
    status: statusFilter,
  });
  const deleteTask = useDeleteTask();

  return (
    <div className="space-y-6">
      {/* Page header */}
      <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Tasks</h1>
          <p className="text-muted-foreground mt-1">
            Manage your tasks and stay organized
          </p>
        </div>
        <Button render={<Link href="/tasks/new" />}>
          <Plus className="mr-2 h-4 w-4" />
          New Task
        </Button>
      </div>

      {/* Filters */}
      <div className="flex flex-col gap-3 sm:flex-row">
        <div className="relative flex-1">
          <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
          <Input
            placeholder="Search tasks..."
            className="pl-9"
            value={search}
            onChange={(e) => setSearch(e.target.value)}
          />
        </div>
        <div className="flex gap-2">
          {['all', 'todo', 'in_progress', 'done'].map((status) => (
            <Button
              key={status}
              variant={statusFilter === status || (!statusFilter && status === 'all') ? 'default' : 'outline'}
              size="sm"
              onClick={() => setStatusFilter(status === 'all' ? undefined : status)}
            >
              {status === 'all' ? 'All' : status.replace('_', ' ')}
            </Button>
          ))}
        </div>
      </div>

      {/* Loading state */}
      {isLoading && (
        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
          {Array.from({ length: 6 }).map((_, i) => (
            <Card key={i}>
              <CardHeader>
                <Skeleton className="h-5 w-3/4" />
                <Skeleton className="h-4 w-1/2 mt-2" />
              </CardHeader>
              <CardContent>
                <Skeleton className="h-4 w-full" />
                <Skeleton className="h-4 w-2/3 mt-2" />
              </CardContent>
              <CardFooter>
                <Skeleton className="h-6 w-16" />
              </CardFooter>
            </Card>
          ))}
        </div>
      )}

      {/* Error state */}
      {isError && (
        <Card className="border-destructive">
          <CardContent className="flex flex-col items-center justify-center py-12">
            <AlertCircle className="h-12 w-12 text-destructive mb-4" />
            <h3 className="text-lg font-semibold mb-2">Failed to load tasks</h3>
            <p className="text-muted-foreground mb-4 text-center max-w-md">
              {error?.message || 'An unexpected error occurred. Please try again.'}
            </p>
            <Button variant="outline" onClick={() => refetch()}>
              <RefreshCw className="mr-2 h-4 w-4" />
              Retry
            </Button>
          </CardContent>
        </Card>
      )}

      {/* Empty state */}
      {!isLoading && !isError && data?.tasks.length === 0 && (
        <Card>
          <CardContent className="flex flex-col items-center justify-center py-12">
            <div className="rounded-full bg-muted p-4 mb-4">
              <Search className="h-8 w-8 text-muted-foreground" />
            </div>
            <h3 className="text-lg font-semibold mb-2">No tasks found</h3>
            <p className="text-muted-foreground mb-4 text-center max-w-md">
              {search || statusFilter
                ? 'No tasks match your current filters. Try adjusting your search.'
                : 'Get started by creating your first task.'}
            </p>
            {!search && !statusFilter && (
              <Button render={<Link href="/tasks/new" />}>
                <Plus className="mr-2 h-4 w-4" />
                Create your first task
              </Button>
            )}
          </CardContent>
        </Card>
      )}

      {/* Task grid */}
      {!isLoading && !isError && data?.tasks && data.tasks.length > 0 && (
        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
          {data.tasks.map((task: Task) => (
            <Card
              key={task.id}
              className={cn(
                'group transition-shadow hover:shadow-md',
                task.status === 'done' && 'opacity-70'
              )}
            >
              <CardHeader className="pb-2">
                <div className="flex items-start justify-between gap-2">
                  <CardTitle className="text-base line-clamp-2">
                    <Link
                      href={`/tasks/${task.id}`}
                      className="hover:underline"
                    >
                      {task.title}
                    </Link>
                  </CardTitle>
                  <DropdownMenu>
                    <DropdownMenuTrigger
                      render={
                        <Button
                          variant="ghost"
                          size="icon"
                          className="h-8 w-8 opacity-0 group-hover:opacity-100 transition-opacity"
                        />
                      }
                    >
                      <span className="sr-only">Open menu</span>
                      <span className="text-lg leading-none">⋯</span>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent align="end">
                      <DropdownMenuItem
                        render={<Link href={`/tasks/${task.id}`} />}
                      >
                        View
                      </DropdownMenuItem>
                      <DropdownMenuItem
                        render={<Link href={`/tasks/${task.id}/edit`} />}
                      >
                        Edit
                      </DropdownMenuItem>
                      <DropdownMenuItem
                        className="text-destructive"
                        onClick={() => {
                          if (confirm('Are you sure you want to delete this task?')) {
                            deleteTask.mutate(task.id);
                          }
                        }}
                      >
                        Delete
                      </DropdownMenuItem>
                    </DropdownMenuContent>
                  </DropdownMenu>
                </div>
              </CardHeader>
              <CardContent className="pb-2">
                {task.description && (
                  <p className="text-sm text-muted-foreground line-clamp-2">
                    {task.description}
                  </p>
                )}
              </CardContent>
              <CardFooter className="flex items-center gap-2 pt-0">
                <Badge
                  variant="secondary"
                  className={cn('text-xs', statusColors[task.status])}
                >
                  {task.status.replace('_', ' ')}
                </Badge>
                <Badge
                  variant="secondary"
                  className={cn('text-xs', priorityColors[task.priority])}
                >
                  {task.priority}
                </Badge>
                {task.dueDate && (
                  <span className="ml-auto text-xs text-muted-foreground">
                    {formatDate(task.dueDate)}
                  </span>
                )}
              </CardFooter>
            </Card>
          ))}
        </div>
      )}
    </div>
  );
}
