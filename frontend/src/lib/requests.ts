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
