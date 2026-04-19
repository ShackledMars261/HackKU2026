import { getPostsByLocation } from '@/api/post/getPostsByLocation';
import type { PageServerLoad } from './$types';
import { getLocation } from '@/api/location/getLocation';

export const load: PageServerLoad = async ({ locals, params }) => {
	const { slug } = params;

	const token: string = locals.sessionToken || '';

	const location = await getLocation(token, slug);
	console.log(JSON.stringify(location));

	const posts = await getPostsByLocation(token, slug);
	console.log(JSON.stringify(posts));

	return { location, posts };
};
