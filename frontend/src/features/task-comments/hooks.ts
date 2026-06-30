'use client';

import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { toast } from 'sonner';
import * as commentsApi from './api';
import { taskCommentKeys, taskCommentsQueryOptions } from './queries';
import type { CreateTaskCommentInput } from './types';

export function useTaskComments(taskId: string) {
  return useQuery(taskCommentsQueryOptions(taskId));
}

export function useCreateTaskComment(taskId: string) {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: CreateTaskCommentInput) => commentsApi.createTaskComment(taskId, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: taskCommentKeys.list(taskId) });
      toast.success('Comment posted');
    },
    onError: (error: Error) => {
      toast.error(error.message || 'Failed to post comment');
    },
  });
}
