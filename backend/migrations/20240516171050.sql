-- Create "users" table
CREATE TABLE "public"."users" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  "mmr" bigint NULL,
  "mu" numeric NULL,
  "sigma" numeric NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_users_name" UNIQUE ("name")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "public"."users" ("deleted_at");
-- Create "teams" table
CREATE TABLE "public"."teams" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "user_one_id" bigint NULL,
  "user_two_id" bigint NULL,
  "score" bigint NULL,
  "winner" boolean NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_teams_user_one" FOREIGN KEY ("user_one_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_teams_user_two" FOREIGN KEY ("user_two_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_teams_deleted_at" to table: "teams"
CREATE INDEX "idx_teams_deleted_at" ON "public"."teams" ("deleted_at");
-- Create "matches" table
CREATE TABLE "public"."matches" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "team_one_id" bigint NULL,
  "team_two_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_matches_team_one" FOREIGN KEY ("team_one_id") REFERENCES "public"."teams" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_matches_team_two" FOREIGN KEY ("team_two_id") REFERENCES "public"."teams" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_matches_deleted_at" to table: "matches"
CREATE INDEX "idx_matches_deleted_at" ON "public"."matches" ("deleted_at");
