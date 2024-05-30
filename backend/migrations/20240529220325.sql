-- Modify "mmr_calculations" table
ALTER TABLE "mmr_calculations" DROP CONSTRAINT "fk_mmr_calculations_match", ADD
 CONSTRAINT "fk_matches_mmr_calculations" FOREIGN KEY ("match_id") REFERENCES "matches" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
