import { getPostsByLocation } from '@/api/post/getPostsByLocation';
import type { PageServerLoad } from './$types';
import { getLocation } from '@/api/location/getLocation';
import { getStatusForLocation } from '@/api/status/getStatus';

export const load: PageServerLoad = async ({ locals, params }) => {
	const { slug } = params;

	const token: string = locals.sessionToken || '';

	const location = await getLocation(token, slug);
	console.log(JSON.stringify(location));

	const posts = await getPostsByLocation(token, slug);
	console.log(JSON.stringify(posts));

	const status = await getStatusForLocation(token, slug, '1h');
	console.log(JSON.stringify(status));

	return { location, posts, status };
};
