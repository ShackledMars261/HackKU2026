import { error } from '@sveltejs/kit';
import type { User } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals, params, fetch }) => {
	const { slug } = params;

	const res = await fetch(`http://${process.env.BACKEND_URL}/user/${slug}`, {
		headers: {
			Authorization: `Bearer ${locals.sessionToken}`
		}
	});

	if (res.status === 404) {
		error(404, 'User not found');
	}

	if (!res.ok) {
		error(500, 'Failed to fetch user');
	}

	const user: User = await res.json();
	return { user };
};
