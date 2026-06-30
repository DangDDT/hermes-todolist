import { queryOptions } from '@tanstack/react-query';
import * as tasksApi from './api';

export const taskKeys = {
  all: ['tasks'] as const,
  lists: () => [...taskKeys.all, 'list'] as const,
  list: (filters?: Record<string, string | number | undefined>) =>
    [...taskKeys.lists(), filters] as const,
  details: () => [...taskKeys.all, 'detail'] as const,
  detail: (id: string) => [...taskKeys.details(), id] as const,
};

export function tasksQueryOptions(filters?: {
  page?: number;
  limit?: number;
  status?: string;
  priority?: string;
  search?: string;
}) {
  return queryOptions({
    queryKey: taskKeys.list(filters),
    queryFn: () => tasksApi.getTasks(filters),
  });
}

export function taskQueryOptions(id: string) {
  return queryOptions({
    queryKey: taskKeys.detail(id),
    queryFn: () => tasksApi.getTask(id),
  });
}
