import { queryOptions } from '@tanstack/react-query';
import * as commentsApi from './api';

export const taskCommentKeys = {
  all: ['task-comments'] as const,
  list: (taskId: string) => [...taskCommentKeys.all, taskId] as const,
};

export function taskCommentsQueryOptions(taskId: string) {
  return queryOptions({
    queryKey: taskCommentKeys.list(taskId),
    queryFn: () => commentsApi.getTaskComments(taskId),
  });
}
