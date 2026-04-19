import type { RequestHandler } from '@sveltejs/kit';
import { createStatusReport } from '@/api/status/createStatusReport';
import { json, error } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ request, locals }) => {
	const token = locals.sessionToken;
	const { location, busyness } = await request.json();

	console.log(location);
	console.log(busyness);
	console.log(token);

	const statusReport = await createStatusReport(token, {
		location_id: location,
		busyness: busyness,
		loudness: busyness
	});

	return json(statusReport);
};
