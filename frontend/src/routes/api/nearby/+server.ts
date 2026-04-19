import type { RequestHandler } from './$types';
import { getNearbyLocations } from '@/api/location/getNearbyLocations';
import { json, error } from '@sveltejs/kit';

const ONE_MILE_IN_METERS = 1609;

export const POST: RequestHandler = async ({ request, locals }) => {
	const token = locals.sessionToken; // adjust to match your locals shape
	const { latitude, longitude } = await request.json();

	if (!latitude || !longitude) {
		throw error(400, 'Missing coordinates');
	}

	const locations = await getNearbyLocations(token, {
		latitude,
		longitude,
		radius: ONE_MILE_IN_METERS
	});

	return json(locations ?? []);
};
