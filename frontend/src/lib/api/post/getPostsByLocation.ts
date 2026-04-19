import type { Post } from '@/types';

export async function getPostsByLocation(
	token: string,
	locationId: string
): Promise<Post[] | void> {
	try {
		const resp = await fetch(`http://${process.env.BACKEND_URL}:8080/posts/${locationId}`, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${token}`
			}
		});

		if (!resp.ok) {
			throw new Error(`Error: ${resp.status}`);
		}

		const data: Post[] = await resp.json();
		console.log(JSON.stringify(data));
		return data;
	} catch (error) {
		console.error('Fetch error:', error);
	}
}
