package repos

import (
	"mmr/backend/db/models"

	"gorm.io/gorm"
)

type ISeasonRepository interface {
	CurrentSeason() (*models.Season, error)
	CreateSeason(season *models.Season) (*models.Season, error)
}

type SeasonRepository struct {
	db *gorm.DB
}

func NewSeasonRepository(db *gorm.DB) ISeasonRepository {
	return &SeasonRepository{db}
}

func (sr *SeasonRepository) CreateSeason(season *models.Season) (*models.Season, error) {
	if err := sr.db.Create(season).Error; err != nil {
		return nil, err
	}
	return season, nil
}

func (sr *SeasonRepository) CurrentSeason() (*models.Season, error) {
	season := &models.Season{}
	if err := sr.db.Order("created_at desc").First(season).Error; err != nil {
		return nil, err
	}
	return season, nil
}
