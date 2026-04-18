import type { SignInRequest } from '@/requests';
import type { Session } from '@/types';

export async function signIn(body: SignInRequest): Promise<Session | void> {
	try {
		const resp = await fetch(`http://${process.env.BACKEND_URL}:8080/signin`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(body)
		});

		if (!resp.ok) {
			throw new Error(`Error: ${resp.status}`);
		}

		const data: Session = await resp.json();
		return data;
	} catch (error) {
		console.error('Fetch error:', error);
	}
}
