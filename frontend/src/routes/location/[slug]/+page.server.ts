import type { PageServerLoad } from './$types';
import { getLocation } from '@/api/location/getLocation';

export const load: PageServerLoad = async ({ locals, params }) => {
	const { slug } = params;

	const token: string = locals.sessionToken || '';

	const loc = await getLocation(token, slug);

	return { loc };
};
