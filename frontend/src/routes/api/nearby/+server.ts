import type { RequestHandler } from './$types';
import { getNearbyLocations } from '@/api/location/getNearbyLocations';
import { json, error } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ request, locals }) => {
	const token = locals.sessionToken; // adjust to match your locals shape
	const { latitude, longitude } = await request.json();

	if (!latitude || !longitude) {
		throw error(400, 'Missing coordinates');
	}

	const locations = await getNearbyLocations(token, {
		latitude: latitude,
		longitude: longitude,
		radius: 1
	});

	return json(locations ?? []);
};
