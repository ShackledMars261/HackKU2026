import { redirect } from "@sveltejs/kit";
import type { Handle } from "@sveltejs/kit";

const PROTECTED_ROUTES = ['/profile'];

export const handle: Handle = async ({ event, resolve }) => {
    const isProtected = PROTECTED_ROUTES.some(route => 
    event.url.pathname.startsWith(route)
  );

  if (!isProtected) {
    return resolve(event);
  }

  const sessionToken = event.cookies.get('session');

  if (!sessionToken) {
    redirect(303, '/login');
  }

  const res = await event.fetch(`${process.env.BACKEND_URL}/auth/session`, {
    headers: {
      Authorization: `Bearer ${sessionToken}`
    }
  });

  if (!res.ok) {
    redirect(303, '/login');
  }

  const user = await res.json();

  event.locals.user = user;

  return resolve(event);
};