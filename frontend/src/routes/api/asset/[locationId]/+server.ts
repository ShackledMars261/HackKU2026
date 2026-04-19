import type { RequestHandler } from './$types';
import { error } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ request, locals, params }) => {
	const token: string = locals.sessionToken || '';
	const filename = request.headers.get('X-Filename') || 'upload';
	const { locationId } = params;

	const buffer = await request.arrayBuffer();

	const resp = await fetch(`http://${process.env.BACKEND_URL}:8080/asset/${locationId}`, {
		method: 'POST',
		headers: {
			Authorization: `Bearer ${token}`,
			'Content-Type': 'application/octet-stream',
			'X-Filename': filename
		},
		body: buffer
	});

	if (!resp.ok) {
		throw error(resp.status, await resp.text());
	}

	return new Response(resp.body, {
		headers: { 'Content-Type': 'application/json' }
	});
};
