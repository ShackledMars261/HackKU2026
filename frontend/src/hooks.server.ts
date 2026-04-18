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

  const res = await event.fetch(`${process.env.BACKEND_URL}/session/${sessionToken}`, {
    headers: {
      Authorization: `Bearer ${sessionToken}`
    }
  });

  if (!res.ok) {
    redirect(303, '/login');
  }

  const resp = await res.json();

  const userRes = await event.fetch(`${process.env.BACKEND_URL}/profile/${resp.userId}`, {
    headers: {
      Authorization: `Bearer ${sessionToken}`
    }
  });

  const user = await userRes.json();

  event.locals.user = user;

  return resolve(event);
};