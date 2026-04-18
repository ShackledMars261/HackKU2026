import { redirect, type Actions } from '@sveltejs/kit';
import type { CreateLocationRequest } from '@/requests';
import { createLocation } from '@/api/location/createLocation';

export const actions: Actions = {
	submitLocation: async ({ cookies, request }) => {
		const data = await request.formData();
		const name: string = data.get('name')?.toString() || '';
		const longitude: number = Number(data.get('longitude')?.toString());
		const latitude: number = Number(data.get('latitude')?.toString());
		const body: CreateLocationRequest = { longitude, latitude, name };
		const token: string = cookies.get('session') || '';
		const location = await createLocation(token, body);
		if (!location) {
			console.log('error');
		} else {
			console.log(JSON.stringify(location));
			redirect(303, '/');
		}
	}
};
