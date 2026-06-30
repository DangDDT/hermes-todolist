export interface TaskComment {
  id: string;
  taskId: string;
  authorId: string;
  authorName: string;
  body: string;
  createdAt: string;
}

export interface TaskCommentsResponse {
  comments: TaskComment[];
}

export interface CreateTaskCommentInput {
  body: string;
}
