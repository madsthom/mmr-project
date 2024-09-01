package services

import (
	database "mmr/backend/db"
	"mmr/backend/db/repos"
)

type SeasonService struct{}

func (ss SeasonService) CurrentSeasonID() uint {
	seasonRepo := repos.NewSeasonRepository(database.DB)
	season, err := seasonRepo.CurrentSeason()
	if err != nil {
		panic("Failed to get a current season. Please create one")
	}
	return season.ID
}
