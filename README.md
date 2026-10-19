# REST API - Film Management
 
A RESTful API for managing film data, built with Go, Gin, GORM, and PostgreSQL using Clean Architecture principles.
 
## Tech Stack
 
- **Language:** Go (Golang)
- **Framework:** Gin
- **ORM:** GORM
- **Database:** PostgreSQL
- **Containerization:** Docker & Docker Compose
## Architecture
 
This project uses a layered architecture:
 
`Handler` ➡️ `Service` ➡️ `Repository` ➡️ `PostgreSQL`
 
## How to Run
 
### Using Docker (Recommended)
 
1. Clone this repository.
2. Make sure Docker is running.
3. Run the application:
```bash
   docker compose up --build -d
```
 
4. The API will be available at <http://localhost:8080>.
### Running Locally
 
1. Set up `.env` based on `.env.example`.
2. Run PostgreSQL (via Docker):
```bash
   docker compose up -d postgres
```
 
3. Run the app:
```bash
   go run cmd/api/main.go
```
 
## API Endpoints
 
| Method | Endpoint         | Description        |
| ------ | ---------------- | ------------------ |
| POST   | `/api/films`     | Create a new film  |
| GET    | `/api/films`     | Get all films      |
| GET    | `/api/films/:id` | Get film by ID     |
| PUT    | `/api/films/:id` | Update a film      |
| DELETE | `/api/films/:id` | Delete a film      |
 
### Example Request (`POST /api/films`)
 
```json
{
  "judul": "Interstellar",
  "sutradara": "Christopher Nolan",
  "tahun": 2014,
  "genre": "Sci-Fi",
  "rating": 9.0
}
```
 
