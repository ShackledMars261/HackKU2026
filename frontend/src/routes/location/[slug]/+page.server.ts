import { getPostsByLocation } from '@/api/post/getPostsByLocation';
import type { PageServerLoad } from './$types';
import { getLocation } from '@/api/location/getLocation';
import { getStatusForLocation } from '@/api/status/getStatus';
import { getAllAssetUrlsForLocation } from '@/api/assets/getAllAssetUrls';

export const load: PageServerLoad = async ({ locals, params }) => {
	const { slug } = params;
	const token: string = locals.sessionToken || '';

	const [location, posts, status, assetUrls] = await Promise.allSettled([
		getLocation(token, slug),
		getPostsByLocation(token, slug),
		getStatusForLocation(token, slug, '1h'),
		getAllAssetUrlsForLocation(token, slug)
	]);

	return {
		location: location.status === 'fulfilled' ? location.value : null,
		posts: posts.status === 'fulfilled' ? (posts.value ?? []) : [],
		status: status.status === 'fulfilled' ? (status.value ?? null) : null,
		assetUrls: assetUrls.status === 'fulfilled' ? (assetUrls.value ?? { paths: [] }) : { paths: [] }
	};
};
