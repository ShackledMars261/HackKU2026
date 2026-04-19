import type { CreateLocationRequest } from '@/requests';
import type { Location } from '@/types';

export async function createLocation(
	token: string,
	body: CreateLocationRequest
): Promise<Location | void> {
	try {
		const resp = await fetch(`http://${process.env.BACKEND_URL}:8080/location`, {
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

		const data: Location = await resp.json();
		return data;
	} catch (error) {
		console.error('Fetch error:', error);
	}
}
