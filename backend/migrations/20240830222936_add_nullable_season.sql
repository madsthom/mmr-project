-- Create "seasons" table
CREATE TABLE "seasons" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_seasons_deleted_at" to table: "seasons"
CREATE INDEX "idx_seasons_deleted_at" ON "seasons" ("deleted_at");
-- Modify "matches" table
ALTER TABLE "matches" ADD COLUMN "season_id" bigint NULL, ADD
 CONSTRAINT "fk_matches_season" FOREIGN KEY ("season_id") REFERENCES "seasons" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
