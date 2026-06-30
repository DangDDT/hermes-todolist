'use client';

import { useState } from 'react';
import Link from 'next/link';
import { Plus, Search, AlertCircle, RefreshCw, ArrowUpDown } from 'lucide-react';
import { useTasks, useDeleteTask } from '@/features/tasks/hooks';
import type { Task } from '@/features/tasks/types';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import { Skeleton } from '@/components/ui/skeleton';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import { cn, formatDate } from '@/lib/utils';

const PAGE_SIZE = 12;

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

const sortOptions = [
  { value: 'createdAt:desc', label: 'Newest First' },
  { value: 'createdAt:asc', label: 'Oldest First' },
  { value: 'dueDate:asc', label: 'Due Date (Earliest)' },
  { value: 'dueDate:desc', label: 'Due Date (Latest)' },
  { value: 'priority:asc', label: 'Priority (Low→High)' },
  { value: 'priority:desc', label: 'Priority (High→Low)' },
];

export default function TasksPage() {
  const [search, setSearch] = useState('');
  const [statusFilter, setStatusFilter] = useState<string | undefined>();
  const [priorityFilter, setPriorityFilter] = useState<string | undefined>();
  const [sort, setSort] = useState('createdAt:desc');
  const [page, setPage] = useState(1);
  const [deleteTarget, setDeleteTarget] = useState<Task | null>(null);

  const [sortField, sortOrder] = sort.split(':') as [string, string];

  const { data, isLoading, isError, error, refetch } = useTasks({
    page,
    limit: PAGE_SIZE,
    search: search || undefined,
    status: statusFilter,
    priority: priorityFilter,
    sortBy: sortField,
    sortOrder,
  });

  const deleteTask = useDeleteTask();

  const totalPages = data ? Math.ceil(data.total / PAGE_SIZE) : 0;

  function handleDelete() {
    if (deleteTarget) {
      deleteTask.mutate(deleteTarget.id);
      setDeleteTarget(null);
    }
  }

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
            onChange={(e) => {
              setSearch(e.target.value);
              setPage(1);
            }}
          />
        </div>
        <div className="flex gap-2 flex-wrap">
          {/* Status filter */}
          <Select
            value={statusFilter || 'all'}
            onValueChange={(value) => {
              setStatusFilter(!value || value === 'all' ? undefined : value);
              setPage(1);
            }}
          >
            <SelectTrigger className="w-[130px]">
              <SelectValue placeholder="Status" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="all">All Status</SelectItem>
              <SelectItem value="todo">To Do</SelectItem>
              <SelectItem value="in_progress">In Progress</SelectItem>
              <SelectItem value="done">Done</SelectItem>
            </SelectContent>
          </Select>

          {/* Priority filter */}
          <Select
            value={priorityFilter || 'all'}
            onValueChange={(value) => {
              setPriorityFilter(!value || value === 'all' ? undefined : value);
              setPage(1);
            }}
          >
            <SelectTrigger className="w-[130px]">
              <SelectValue placeholder="Priority" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="all">All Priority</SelectItem>
              <SelectItem value="low">Low</SelectItem>
              <SelectItem value="medium">Medium</SelectItem>
              <SelectItem value="high">High</SelectItem>
            </SelectContent>
          </Select>

          {/* Sort */}
          <Select
            value={sort}
            onValueChange={(value) => {
              setSort(value ?? sort);
              setPage(1);
            }}
          >
            <SelectTrigger className="w-[180px]">
              <ArrowUpDown className="mr-1 h-3.5 w-3.5 text-muted-foreground" />
              <SelectValue placeholder="Sort by" />
            </SelectTrigger>
            <SelectContent>
              {sortOptions.map((opt) => (
                <SelectItem key={opt.value} value={opt.value}>
                  {opt.label}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
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
              {search || statusFilter || priorityFilter
                ? 'No tasks match your current filters. Try adjusting your search.'
                : 'Get started by creating your first task.'}
            </p>
            {!search && !statusFilter && !priorityFilter && (
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
        <>
          <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
            {data.tasks.map((task: Task) => (
              <Card
                key={task.id}
                className={cn(
                  'group transition-shadow hover:shadow-md cursor-pointer',
                  task.status === 'done' && 'opacity-70'
                )}
                onClick={() => {
                  // Navigate to detail unless clicking the dropdown
                }}
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
                            onClick={(e: React.MouseEvent) => e.stopPropagation()}
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
                          onClick={(e) => {
                            e.stopPropagation();
                            setDeleteTarget(task);
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

          {/* Pagination */}
          {totalPages > 1 && (
            <div className="flex items-center justify-between pt-2">
              <p className="text-sm text-muted-foreground">
                Page {page} of {totalPages} ({data.total} tasks)
              </p>
              <div className="flex gap-2">
                <Button
                  variant="outline"
                  size="sm"
                  disabled={page <= 1}
                  onClick={() => setPage((p) => Math.max(1, p - 1))}
                >
                  Previous
                </Button>
                <Button
                  variant="outline"
                  size="sm"
                  disabled={page >= totalPages}
                  onClick={() => setPage((p) => p + 1)}
                >
                  Next
                </Button>
              </div>
            </div>
          )}
        </>
      )}

      {/* Delete confirmation dialog */}
      <Dialog open={!!deleteTarget} onOpenChange={(open) => !open && setDeleteTarget(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Delete Task</DialogTitle>
            <DialogDescription>
              Are you sure you want to delete &ldquo;{deleteTarget?.title}&rdquo;? This action cannot be undone.
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
