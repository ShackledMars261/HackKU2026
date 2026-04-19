import type { RequestHandler } from './$types';
import { json, error } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ request, locals }) => {
	const token: string = locals.sessionToken || '';
	const body = await request.json();

	const resp = await fetch(`http://${process.env.BACKEND_URL}:8080/posts/nopic`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${token}`
		},
		body: JSON.stringify(body)
	});

	if (!resp.ok) {
		throw error(resp.status, await resp.text());
	}

	return json(await resp.json());
};
