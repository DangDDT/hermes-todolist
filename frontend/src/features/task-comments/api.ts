import { apiClient } from '@/lib/api-client';
import type { CreateTaskCommentInput, TaskCommentsResponse } from './types';

export async function getTaskComments(taskId: string): Promise<TaskCommentsResponse> {
  return apiClient<TaskCommentsResponse>(`/tasks/${taskId}/comments`);
}

export async function createTaskComment(
  taskId: string,
  data: CreateTaskCommentInput
): Promise<{ comment: TaskCommentsResponse['comments'][number] }> {
  return apiClient<{ comment: TaskCommentsResponse['comments'][number] }>(`/tasks/${taskId}/comments`, {
    method: 'POST',
    body: JSON.stringify(data),
  });
}
