export type User = {
	id: string;
	username: string;
	password: string;
};

export type Location = {
	id: string;
	latitude: number;
	longitude: number;
	name: string;
	overallRating: number;
};

export type Post = {
	userId: string;
	locationId: string;
	rating: number;
	description: string;
	photoFileIds: string[];
};

export type Session = {
	id: string;
	userId: string;
};
