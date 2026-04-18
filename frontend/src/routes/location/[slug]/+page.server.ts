import { error } from '@sveltejs/kit';
import type { Load } from '@sveltejs/kit';
import type { Location } from '$lib/types';

export const load: Load = async ({ params, fetch }) => {
  const { slug } = params;

  const res = await fetch(`${process.env.BACKEND_URL}/location/${slug}`);

  if (res.status === 404) {
    error(404, 'Location not found');
  }

  if (!res.ok) {
    error(500, 'Failed to fetch location');
  }

  const location: Location = await res.json();
  return { location };
};