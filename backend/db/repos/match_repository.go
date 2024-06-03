package repos

import (
	"gorm.io/gorm/clause"
	"mmr/backend/db/models"

	"gorm.io/gorm"
)

type IMatchRepository interface {
	CreateMatch(match *models.Match) (*models.Match, error)
	ListMatches(limit int, offset int, orderBy *clause.OrderByColumn, includeMmrCalculations bool) ([]*models.Match, error)
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

func (mr *MatchRepository) ListMatches(limit int, offset int, orderBy *clause.OrderByColumn, includeMmrCalculations bool) ([]*models.Match, error) {
	var matches []*models.Match

	if orderBy == nil {
		orderBy = &clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: false}
	}

	query := mr.db.Model(&models.Match{}).
		Preload("TeamOne.UserOne").
		Preload("TeamOne.UserTwo").
		Preload("TeamTwo.UserOne").
		Preload("TeamTwo.UserTwo")

	if includeMmrCalculations {
		query = query.Preload("MMRCalculations")
	}

	err := query.
		Order(*orderBy).
		Limit(limit).
		Offset(offset).
		Find(&matches).Error

	if err != nil {
		return nil, err
	}

	return matches, nil
}

func (mr *MatchRepository) ClearMMRCalculations() {
	mr.db.Exec("TRUNCATE TABLE mmr_calculations RESTART IDENTITY;")
}
