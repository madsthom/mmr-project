-- Insert a new season and retrieve its ID
WITH new_season AS (
    INSERT INTO seasons (created_at, updated_at)
    VALUES (NOW(), NOW())
    RETURNING id
)
-- Update all matches that currently have no season to have the new season ID
UPDATE matches
SET season_id = (SELECT id FROM new_season)
WHERE season_id IS NULL;