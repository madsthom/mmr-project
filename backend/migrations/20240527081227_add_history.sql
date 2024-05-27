-- Create "player_histories" table
CREATE TABLE "public"."player_histories" (
                                  "id" bigserial NOT NULL,
                                  "created_at" timestamptz NULL,
                                  "updated_at" timestamptz NULL,
                                  "deleted_at" timestamptz NULL,
                                  "user_id" bigint NULL,
                                  "mmr" bigint NULL,
                                  "mu" numeric NULL,
                                  "sigma" numeric NULL,
                                  "match_id" bigint NULL,
                                  PRIMARY KEY ("id"),
                                  CONSTRAINT "fk_player_history_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
                                  CONSTRAINT "fk_player_history_match" FOREIGN KEY ("match_id") REFERENCES "public"."matches" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_player_histories_deleted_at" to table: "player_histories"
CREATE INDEX "idx_player_histories_deleted_at" ON "public"."player_histories" ("deleted_at");
-- Create "mmr_calculations" table
CREATE TABLE "public"."mmr_calculations" (
                                             "id" bigserial NOT NULL,
                                             "created_at" timestamptz NULL,
                                             "updated_at" timestamptz NULL,
                                             "deleted_at" timestamptz NULL,
                                             "match_id" bigint NULL,
                                             "team_one_player_one_mmr_delta" bigint NULL,
                                             "team_one_player_two_mmr_delta" bigint NULL,
                                             "team_two_player_one_mmr_delta" bigint NULL,
                                             "team_two_player_two_mmr_delta" bigint NULL,
                                             PRIMARY KEY ("id"),
                                             CONSTRAINT "fk_mmr_calculations_match" FOREIGN KEY ("match_id") REFERENCES "public"."matches" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_mmr_calculations_deleted_at" to table: "mmr_calculations"
CREATE INDEX "idx_mmr_calculations_deleted_at" ON "public"."mmr_calculations" ("deleted_at");
