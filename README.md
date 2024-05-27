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

First start a local postgres db. You can use the docker-compose file in the backend directory.

```bash
go run main.go
```

### Testing - Backend

Run `go test ./test/...` in the backend folder to run tests

### Migrate db

We use [Atlas](https://atlasgo.io/) for database migrations.

You can install it with in various ways ([found here](https://atlasgo.io/getting-started#installation)). For example, with brew:

```bash
brew install ariga/tap/atlas
```

To apply the migrations run:

```bash
atlas migrate apply --env gorm --revisions-schema public --allow-dirty`
```

#### New migration

In order to add a new migration, you can run:

```bash
atlas migrate diff <name_of_migration> --env gorm
```

This will create a new migration file in the `migrations` folder with the name `YYYYMMddHHmmss<name_of_migration>.sql` based on the changes you've made to the GORM models.
