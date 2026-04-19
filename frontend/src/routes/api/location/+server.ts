import type { RequestHandler } from './$types';
import { createLocation } from '@/api/location/createLocation';
import { json, error } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ request, locals }) => {
	const token: string = locals.sessionToken || '';
	const body = await request.json();

	if (!body.name || body.latitude == null || body.longitude == null) {
		throw error(400, 'Missing required fields');
	}

	const location = await createLocation(token, body);

	if (!location) {
		throw error(500, 'Failed to create location');
	}

	return json(location);
};
