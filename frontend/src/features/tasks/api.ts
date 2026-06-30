import { apiClient } from '@/lib/api-client';
import type { Task, CreateTaskInput, UpdateTaskInput, TaskListResponse } from './types';

export async function getTasks(params?: {
  page?: number;
  limit?: number;
  status?: string;
  priority?: string;
  search?: string;
}): Promise<TaskListResponse> {
  const searchParams = new URLSearchParams();
  if (params?.page) searchParams.set('page', String(params.page));
  if (params?.limit) searchParams.set('limit', String(params.limit));
  if (params?.status) searchParams.set('status', params.status);
  if (params?.priority) searchParams.set('priority', params.priority);
  if (params?.search) searchParams.set('search', params.search);

  const query = searchParams.toString();
  return apiClient<TaskListResponse>(`/tasks${query ? `?${query}` : ''}`);
}

export async function getTask(id: string): Promise<{ task: Task }> {
  return apiClient<{ task: Task }>(`/tasks/${id}`);
}

export async function createTask(data: CreateTaskInput): Promise<{ task: Task }> {
  return apiClient<{ task: Task }>('/tasks', {
    method: 'POST',
    body: JSON.stringify(data),
  });
}

export async function updateTask(
  id: string,
  data: UpdateTaskInput
): Promise<{ task: Task }> {
  return apiClient<{ task: Task }>(`/tasks/${id}`, {
    method: 'PATCH',
    body: JSON.stringify(data),
  });
}

export async function deleteTask(id: string): Promise<void> {
  await apiClient<void>(`/tasks/${id}`, {
    method: 'DELETE',
  });
}
