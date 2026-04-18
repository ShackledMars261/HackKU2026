import { error } from '@sveltejs/kit';
import type { Location } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals, params, fetch }) => {
	const { slug } = params;

	const res = await fetch(`http://${process.env.BACKEND_URL}:8080/location/${slug}`, {
		headers: {
			Authorization: `Bearer ${locals.sessionToken}`
		}
	});

	if (res.status === 404) {
		error(404, 'Location not found');
	}

	if (!res.ok) {
		error(500, 'Failed to fetch location');
	}

	const location: Location = await res.json();
	return { location };
};
