export type User = {
	id: string;
	username: string;
	password: string;
};

export type GeoJSON = {
	type: string;
	coordinates: number[];
};

export type Location = {
	id: string;
	location: GeoJSON;
	owner: string;
	name: string;
	overallRating: number;
};

export type Post = {
	id: string;
	userId: string;
	locationId: string;
	rating: number;
	description: string;
	createdAt: string;
	photoFileIds: string[];
};

export type Session = {
	id: string;
	userId: string;
};

export type StatusResponse = {
	locationId: string;
	averageBusyness: number;
	averageLoudness: number;
};

export type StatusReport = {
	id: string;
	locationId: string;
	userId: string;
	busyness: number;
	loudness: number;
	createdAt: string;
};
