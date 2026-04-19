import type { RequestHandler } from './$types';
import { error } from '@sveltejs/kit';

export const GET: RequestHandler = async ({ params, locals }) => {
	const token: string = locals.sessionToken || '';

	const resp = await fetch(`http://${process.env.BACKEND_URL}:8080/${params.path}`, {
		headers: { Authorization: `Bearer ${token}` }
	});

	if (!resp.ok) {
		throw error(resp.status);
	}

	return new Response(resp.body, {
		headers: {
			'Content-Type': resp.headers.get('Content-Type') || 'application/octet-stream',
			'Cache-Control': 'public, max-age=31536000, immutable'
		}
	});
};
