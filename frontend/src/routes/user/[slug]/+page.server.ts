import { error } from '@sveltejs/kit';
import type { Load } from '@sveltejs/kit';
import type { User } from '$lib/types';

export const load: Load = async ({ params, fetch }) => {
  const { slug } = params;

  const res = await fetch(`${process.env.BACKEND_URL}/user/${slug}`);

  if (res.status === 404) {
    error(404, 'User not found');
  }

  if (!res.ok) {
    error(500, 'Failed to fetch user');
  }

  const user: User = await res.json();
  return { user };
};