'use client';

import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { useRouter } from 'next/navigation';
import { toast } from 'sonner';
import * as tasksApi from './api';
import { taskKeys, tasksQueryOptions, taskQueryOptions } from './queries';
import type { CreateTaskInput, UpdateTaskInput } from './types';

export function useTasks(filters?: {
  page?: number;
  limit?: number;
  status?: string;
  priority?: string;
  search?: string;
}) {
  return useQuery(tasksQueryOptions(filters));
}

export function useTask(id: string) {
  return useQuery(taskQueryOptions(id));
}

export function useCreateTask() {
  const queryClient = useQueryClient();
  const router = useRouter();

  return useMutation({
    mutationFn: (data: CreateTaskInput) => tasksApi.createTask(data),
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: taskKeys.lists() });
      toast.success('Task created successfully');
      router.push(`/tasks/${data.task.id}`);
    },
    onError: (error: Error) => {
      toast.error(error.message || 'Failed to create task');
    },
  });
}

export function useUpdateTask(id: string) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: UpdateTaskInput) => tasksApi.updateTask(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: taskKeys.detail(id) });
      queryClient.invalidateQueries({ queryKey: taskKeys.lists() });
      toast.success('Task updated successfully');
    },
    onError: (error: Error) => {
      toast.error(error.message || 'Failed to update task');
    },
  });
}

export function useDeleteTask() {
  const queryClient = useQueryClient();
  const router = useRouter();

  return useMutation({
    mutationFn: (id: string) => tasksApi.deleteTask(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: taskKeys.lists() });
      toast.success('Task deleted successfully');
      router.push('/tasks');
    },
    onError: (error: Error) => {
      toast.error(error.message || 'Failed to delete task');
    },
  });
}
