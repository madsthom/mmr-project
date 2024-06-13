# MMR Project

This repository contains the source code for the MMR Project, a web application combining a Go backend using Gin and Swagger with a SvelteKit frontend. The backend structure is loosly based on the repo: [Go Gin Boilerplate](https://github.com/vsouza/go-gin-boilerplate).

## Local development

Since we are using Supabase for authentication you will need to start Supabase locally. You can do this by running the following command ([install CLI by following these steps](https://supabase.com/docs/guides/cli/getting-started)):

```bash
supabase start
```

**Note:** You will need to have Docker installed to run Supabase locally.

Use the following values from the output of the `supabase start` command in the backend folders `.env` file:

- `JWT secret`: `JWT_SECRET`

Use the following values from the output of the `supabase start` command in the frontend folders `.env` file:

- `anon key`: `PUBLIC_SUPABASE_ANON_KEY`
- `API URL`: `PUBLIC_SUPABASE_URL`

You can now visit your local Dashboard at [http://localhost:54323/](http://localhost:54323/).

You can stop the supabase instance by running:

```bash
supabase stop
```

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

### Import data from prod

To import data from the prod database, you can run:

```bash
./scripts/import_data.sh <resource-group-name> <prod-server-name> <database-name> <tenant-id> <subscription-id> <username>
```

Username has to have access to the prod database.
