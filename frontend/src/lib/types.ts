export type User = {
  id: string;
  email: string;
  password: string;
};

export type Location = {
    id: string;
    latitude: number;
    longitude: number;
    name: string;
    overallRating: number;
}

export type Post = {
    userId: string;
    locationId: string;
    rating: number;
    description: string;
    photoFileIds: string[];
}