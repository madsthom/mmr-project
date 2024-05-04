package repos

import (
	"example.com/m/v2/db/models"
	"gorm.io/gorm"
)

type IMatchRepository interface {
    CreateMatch(match *models.Match) error
    ListMatches() ([]*models.Match, error)
}


type MatchRepository struct {
    db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) IMatchRepository {
    return &MatchRepository{db}
}

func (mr *MatchRepository) CreateMatch(match *models.Match) error {
    return mr.db.Create(match).Error
}

func (mr *MatchRepository) ListMatches() ([]*models.Match, error) {
    var matches []*models.Match
    if err := mr.db.Find(&matches).Error; err != nil {
        return nil, err
    }
    return matches, nil
}
