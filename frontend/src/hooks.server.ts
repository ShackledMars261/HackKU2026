import { redirect } from '@sveltejs/kit';
import type { Handle } from '@sveltejs/kit';

const PROTECTED_ROUTES = ['/location', '/dashboard'];

export const handle: Handle = async ({ event, resolve }) => {
	const isProtected = PROTECTED_ROUTES.some((route) => event.url.pathname.startsWith(route));

	if (event.cookies.get('session')) {
		event.locals.isSignedIn = true;
		event.locals.sessionToken = event.cookies.get('session') || '';
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

	const res = await event.fetch(`http://${process.env.BACKEND_URL}:8080/session/${sessionToken}`, {
		headers: {
			Authorization: `Bearer ${sessionToken}`
		}
	});

	if (!res.ok) {
		event.cookies.delete('session', { path: '/' });
		redirect(303, '/auth');
	}

	const resp: { exists: boolean; expired: boolean; userId: string } = await res.json();

	if (!resp.exists || resp.expired) {
		event.cookies.delete('session', { path: '/' });
		redirect(303, '/auth');
	}

	const userRes = await event.fetch(`http://${process.env.BACKEND_URL}:8080/user/${resp.userId}`, {
		headers: {
			Authorization: `Bearer ${sessionToken}`
		}
	});

	const user = await userRes.json();

	event.locals.user = user;
	event.locals.sessionToken = sessionToken;

	return resolve(event);
};
