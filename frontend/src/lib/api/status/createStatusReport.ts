import type { CreateStatusReportRequest } from '@/requests';
import type { StatusReport } from '@/types';

export async function createStatusReport(
	token: string,
	body: CreateStatusReportRequest
): Promise<StatusReport | void> {
	try {
		const resp = await fetch(`http://${process.env.BACKEND_URL}:8080/status`, {
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

		const data: StatusReport = await resp.json();
		return data;
	} catch (error) {
		console.error('Fetch error:', error);
	}
}
