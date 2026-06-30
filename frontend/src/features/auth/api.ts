import { apiClient } from '@/lib/api-client';
import type { AuthResponse, LoginInput, RegisterInput, User } from './types';

export async function login(data: LoginInput): Promise<AuthResponse> {
  return apiClient<AuthResponse>('/auth/login', {
    method: 'POST',
    body: JSON.stringify(data),
  });
}

export async function register(data: RegisterInput): Promise<AuthResponse> {
  const { confirmPassword, ...body } = data;
  return apiClient<AuthResponse>('/auth/register', {
    method: 'POST',
    body: JSON.stringify(body),
  });
}

export async function logout(): Promise<void> {
  await apiClient<void>('/auth/logout', {
    method: 'POST',
  });
}

export async function getMe(): Promise<{ user: User }> {
  return apiClient<{ user: User }>('/auth/me');
}
