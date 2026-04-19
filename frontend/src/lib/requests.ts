export interface SignInRequest {
	username: string;
	password: string;
}

export interface RegisterRequest {
	username: string;
	password: string;
}

export interface CreateLocationRequest {
	longitude: number;
	latitude: number;
	name: string;
}

export interface GetNearbyLocationsRequest {
	longitude: number;
	latitude: number;
	radius: number;
}

export interface CreateStatusReportRequest {
	location_id: string;
	busyness: number;
	loudness: number;
}
