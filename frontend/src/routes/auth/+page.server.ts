import type { RegisterRequest, SignInRequest } from '$lib/requests';
import type { Actions, PageServerLoad } from './$types';
import { signIn } from '$lib/api/auth/signin';
import { redirect } from '@sveltejs/kit';
import { register } from '$lib/api/auth/register';

export const load: PageServerLoad = async ({ cookies }) => {
	if (cookies.get('session')) {
		redirect(303, '/');
	}
};

export const actions: Actions = {
	login: async ({ cookies, request }) => {
		const data = await request.formData();
		const username: string = data.get('username')?.toString() || '';
		const password: string = data.get('password')?.toString() || '';
		const body: SignInRequest = { username, password };
		const session = await signIn(body);
		if (!session) {
			console.log('error');
		} else {
			cookies.set('session', session.id, { path: '/' });
			redirect(303, '/');
		}
	},
	register: async ({ cookies, request }) => {
		const data = await request.formData();
		const username: string = data.get('username')?.toString() || '';
		const password: string = data.get('password')?.toString() || '';
		const body: RegisterRequest = { username, password };
		const session = await register(body);
		if (!session) {
			console.log('error');
		} else {
			cookies.set('session', session.id, { path: '/' });
			redirect(303, '/');
		}
	}
} satisfies Actions;
