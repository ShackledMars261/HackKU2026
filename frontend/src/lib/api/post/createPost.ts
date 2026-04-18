import type { Post } from '@/types';

export async function createPost(token: string, formData: FormData): Promise<Post | void> {
	try {
		const resp = await fetch(`http://${process.env.BACKEND_URL}:8080/post`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${token}`
			},
			body: formData
		});

		if (!resp.ok) {
			throw new Error(`Error: ${resp.status}`);
		}

		const data: Post = await resp.json();
		return data;
	} catch (error) {
		console.error('Fetch error:', error);
	}
}
