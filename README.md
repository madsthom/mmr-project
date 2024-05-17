# MMR Project

This repository contains the source code for the MMR Project, a web application combining a Go backend using Gin and Swagger with a SvelteKit frontend. The backend structure is loosly based on the repo: [Go Gin Boilerplate](https://github.com/vsouza/go-gin-boilerplate).

## Deployment

### Backend Deployment:

The backend is automatically deployed to fly.io on merges to the main branch.

### Frontend Deployment:

The frontend is automatically deployed to Vercel on merges to the main branch.

## Frontend

TBA

## Backend

Go with the Gin framework and Swagger UI with swag / swagger.

### Get started

Navigate to the backend directory.

_Make sure Go is installed..._

Install dependencies:

```bash
go mod tidy
```

Run the backend server:

```bash
go run main.go
```

### Testing - Backend

Run `go test ./test/...` in the backend folder to run tests

### Migrate db

`atlas migrate apply --var url="[DATABASE_URL]" --env gorm --revisions-schema public --allow-dirty`
