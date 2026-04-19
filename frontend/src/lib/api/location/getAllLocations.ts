import type { Location } from '@/types';

export async function getAllLocations(token: string): Promise<Location[] | void> {
	try {
		const resp = await fetch(`http://${process.env.BACKEND_URL}:8080/location/all`, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${token}`
			}
		});

		if (!resp.ok) {
			throw new Error(`Error: ${resp.status}`);
		}

		const data: Location[] = await resp.json();
		return data;
	} catch (error) {
		console.error('Fetch error:', error);
	}
}
