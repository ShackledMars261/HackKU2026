import type { StatusResponse } from '@/types';

export async function getStatusForLocation(
	token: string,
	locationId: string,
	recency: string
): Promise<StatusResponse | void> {
	try {
		const resp = await fetch(
			`http://${process.env.BACKEND_URL}:8080/status/${locationId}?recency=${recency}`,
			{
				method: 'GET',
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${token}`
				}
			}
		);

		if (!resp.ok) {
			throw new Error(`Error: ${resp.status}`);
		}

		const data: StatusResponse = await resp.json();
		console.log(JSON.stringify(data));
		return data;
	} catch (error) {
		console.error('Fetch error:', error);
	}
}
