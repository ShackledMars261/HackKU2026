# SpotDrop
 
Community-powered study spot discovery. See how busy, loud, and review-worthy any spot is before you pack up and head out.
 
Built at HackKU 2026.
 
## What it does
 
SpotDrop lets you find a good place to study without showing up and hoping for the best. Log in, and you'll see study spots near you including libraries, coffee shops, and campus buildings. Each have live community-reported busyness and noise levels. Browse reviews, see photos from real visitors, and leave your own report to help the next person out.
 
## Stack
 
| Layer | Technology |
|---|---|
| Frontend | SvelteKit + TypeScript |
| Backend | Go + go-chi |
| Database | MongoDB (GridFS for photo storage, geo queries for location) |
 
## Project layout
```
./
├── frontend/        # SvelteKit app, frontend webpage
├── backend/         # Go web server, the backend server
├── envs/            # Environment files (you must create these)
├── docs/            # Additional documentation
└── docker-compose.yml
```

## Running with Docker Compose
 
### Environment files
**`envs/.globals.env`** 
```env
MONGO_INITDB_ROOT_USERNAME=*USER*
MONGO_INITDB_ROOT_PASSWORD=*PASS*
MONGO_INITDB_DATABASE=*INITDB*
```
 
**`envs/.backend.env`** 
```env
MONGO_URI=*MONGO_URI*
```
 
**`envs/.frontend.env`**
```env
BACKEND_URL=backend
```
 
**`envs/.db.env`** 
```env
```
 
### Start the app
 
```bash
docker compose up --build
```
 
Once running, open [http://localhost:3000](http://localhost:3000) in your browser.
 