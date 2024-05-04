package repos

import (
	"mmr/backend/db/models"

	"gorm.io/gorm"
)

type IMatchRepository interface {
	CreateMatch(match *models.Match) (*models.Match, error)
	ListMatches() ([]*models.Match, error)
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
	if err := mr.db.Find(&matches).Error; err != nil {
		return nil, err
	}
	return matches, nil
}
