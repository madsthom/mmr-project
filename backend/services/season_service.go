package services

import (
	"errors"
	"gorm.io/gorm"
	database "mmr/backend/db"
	"mmr/backend/db/models"
	"mmr/backend/db/repos"
)

type SeasonService struct{}

func (ss SeasonService) CurrentSeasonID() uint {
	seasonRepo := repos.NewSeasonRepository(database.DB)
	season, err := seasonRepo.CurrentSeason()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newSeason, newErr := seasonRepo.CreateSeason(&models.Season{})
			if newErr != nil {
				panic("Failed to create new season")
			}
			return newSeason.ID
		}
		panic("Failed to get current season")
	}
	return season.ID
}
