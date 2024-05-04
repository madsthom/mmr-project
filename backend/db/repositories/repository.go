// mysql.go
package repositories

import (
    "gorm.io/gorm"
    "example.com/m/v2/db/models"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
    return &UserRepository{db}
}

func (ur *UserRepository) ListUsers() ([]*models.User, error) {
    var users []*models.User
    if err := ur.db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}

type TeamRepository struct {
    db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) ITeamRepository {
    return &TeamRepository{db}
}

func (tr *TeamRepository) CreateTeam(team *models.Team) error {
    return tr.db.Create(team).Error
}

func (tr *TeamRepository) ListTeams() ([]*models.Team, error) {
    var teams []*models.Team
    if err := tr.db.Find(&teams).Error; err != nil {
        return nil, err
    }
    return teams, nil
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
