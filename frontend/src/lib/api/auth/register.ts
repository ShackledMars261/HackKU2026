import type { RegisterRequest } from '@/requests';
import type { User } from '@/types';

export async function register(body: RegisterRequest): Promise<User | void> {
	try {
		const resp = await fetch(`http://${process.env.BACKEND_URL}/signup`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(body)
		});

		if (!resp.ok) {
			throw new Error(`Error: ${resp.status}`);
		}

		const data: User = await resp.json();
		return data;
	} catch (error) {
		console.error('Fetch error:', error);
	}
}
