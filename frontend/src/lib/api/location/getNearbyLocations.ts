import type { GetNearbyLocationsRequest } from '@/requests';
import type { Location } from '@/types';

export async function getNearbyLocations(
	token: string,
	body: GetNearbyLocationsRequest
): Promise<Location[] | void> {
	try {
		console.log(`http://${process.env.BACKEND_URL}:8080/location/nearby`);
		console.log(JSON.stringify(body));
		console.log(`Bearer ${token}`);
		const resp = await fetch(`http://${process.env.BACKEND_URL}:8080/location/nearby`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${token}`
			},
			body: JSON.stringify(body)
		});

		if (!resp.ok) {
			throw new Error(`Error: ${resp.status}`);
		}

		const data: Location[] = await resp.json();
		data.forEach((loc) => {
			console.log(JSON.stringify(loc.location.coordinates));
		});
		return data;
	} catch (error) {
		console.error('Fetch error:', error);
	}
}
