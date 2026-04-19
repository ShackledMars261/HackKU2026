import type { AssetUrls } from '@/types';

export async function getAllAssetUrlsForLocation(
	token: string,
	locationId: string
): Promise<AssetUrls | void> {
	try {
		const resp = await fetch(`http://${process.env.BACKEND_URL}:8080/asset/all/${locationId}`, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${token}`
			}
		});

		if (!resp.ok) {
			throw new Error(`Error: ${resp.status}`);
		}

		const data: AssetUrls = await resp.json();
		console.log(JSON.stringify(data));
		return data;
	} catch (error) {
		console.error('Fetch error:', error);
	}
}
