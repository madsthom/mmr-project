package repos

import (
	"mmr/backend/db/models"

	"gorm.io/gorm"
)

type IMatchRepository interface {
	CreateMatch(match *models.Match) (*models.Match, error)
	ListMatches() ([]*models.Match, error)
	ClearMMRCalculations()
}

type MatchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) IMatchRepository {
	return &MatchRepository{db}
}

func (mr *MatchRepository) CreateMatch(match *models.Match) (*models.Match, error) {
	if err := mr.db.Create(match).Error; err != nil {
		return nil, err
	}
	return match, nil
}

func (mr *MatchRepository) ListMatches() ([]*models.Match, error) {
	var matches []*models.Match

	err := mr.db.Model(&models.Match{}).
		Preload("TeamOne.UserOne").
		Preload("TeamOne.UserTwo").
		Preload("TeamTwo.UserOne").
		Preload("TeamTwo.UserTwo").
		Order("created_at asc").
		Find(&matches).Error

	if err != nil {
		return nil, err
	}

	return matches, nil
}

func (mr *MatchRepository) ClearMMRCalculations() {
	mr.db.Exec("DELETE FROM mmr_calculations")
}
