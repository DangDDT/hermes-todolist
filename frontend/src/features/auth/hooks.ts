'use client';

import { useMutation } from '@tanstack/react-query';
import { useRouter } from 'next/navigation';
import { useAuth } from './store';
import * as authApi from './api';
import type { LoginInput, RegisterInput } from './types';
import { toast } from 'sonner';

export function useLogin() {
  const router = useRouter();
  const { setUser } = useAuth();

  return useMutation({
    mutationFn: (data: LoginInput) => authApi.login(data),
    onSuccess: (data) => {
      setUser(data.user);
      toast.success('Welcome back!');
      router.push('/tasks');
    },
    onError: (error: Error) => {
      toast.error(error.message || 'Login failed');
    },
  });
}

export function useRegister() {
  const router = useRouter();
  const { setUser } = useAuth();

  return useMutation({
    mutationFn: (data: RegisterInput) => authApi.register(data),
    onSuccess: (data) => {
      setUser(data.user);
      toast.success('Account created successfully!');
      router.push('/tasks');
    },
    onError: (error: Error) => {
      toast.error(error.message || 'Registration failed');
    },
  });
}

export function useLogout() {
  const router = useRouter();
  const { clearUser } = useAuth();

  return useMutation({
    mutationFn: () => authApi.logout(),
    onSuccess: () => {
      clearUser();
      toast.success('Logged out successfully');
      router.push('/login');
    },
    onError: (error: Error) => {
      toast.error(error.message || 'Logout failed');
    },
  });
}
