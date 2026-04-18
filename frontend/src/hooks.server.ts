import { redirect } from '@sveltejs/kit';
import type { Handle } from '@sveltejs/kit';

const PROTECTED_ROUTES = ['/location'];

export const handle: Handle = async ({ event, resolve }) => {
	const isProtected = PROTECTED_ROUTES.some((route) => event.url.pathname.startsWith(route));

	if (event.cookies.get('session')) {
		event.locals.isSignedIn = true;
	} else {
		event.locals.isSignedIn = false;
	}

	if (!isProtected) {
		return resolve(event);
	}

	const sessionToken = event.cookies.get('session');

	if (!sessionToken) {
		redirect(303, '/auth');
	}

	const res = await event.fetch(`http://${process.env.BACKEND_URL}/session/${sessionToken}`, {
		headers: {
			Authorization: `Bearer ${sessionToken}`
		}
	});

	if (!res.ok) {
		redirect(303, '/auth');
	}

	const resp = await res.json();

	const userRes = await event.fetch(`http://${process.env.BACKEND_URL}/profile/${resp.userId}`, {
		headers: {
			Authorization: `Bearer ${sessionToken}`
		}
	});

	const user = await userRes.json();

	event.locals.user = user;
	event.locals.sessionToken = sessionToken;

	return resolve(event);
};
