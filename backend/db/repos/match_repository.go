package repos

import (
	"database/sql"
	"mmr/backend/db/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IMatchRepository interface {
	CreateMatch(match *models.Match) (*models.Match, error)
	ListMatches(limit int, offset int, orderBy *clause.OrderByColumn, includeMmrCalculations bool, userId *uint) ([]*models.Match, error)
	ClearMMRCalculations()
	CheckExistingMatch(playerOneID uint, playerTwoID uint, playerThreeID uint, playerFourID uint, teamOneScore int, teamTwoScore int) bool
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

func (mr *MatchRepository) ListMatches(limit int, offset int, orderBy *clause.OrderByColumn, includeMmrCalculations bool, userId *uint) ([]*models.Match, error) {
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

	if userId != nil {
		query.Joins("TeamOne", mr.db.Where((&models.Team{UserOneID: *userId})))
		query.Joins("TeamOne", mr.db.Where((&models.Team{UserTwoID: *userId})))
		query.Joins("TeamTwo", mr.db.Where((&models.Team{UserOneID: *userId})))
		query.Joins("TeamTwo", mr.db.Where((&models.Team{UserTwoID: *userId})))
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

func (mr *MatchRepository) CheckExistingMatch(playerOneID uint, playerTwoID uint, playerThreeID uint, playerFourID uint, teamOneScore int, teamTwoScore int) bool {
	// Check if match is within 10 minutes of another match
	// This is to prevent spamming of matches
	// It does not care about the order of players or teams

	var count int64

	mr.db.Model(&models.Match{}).
		Joins("JOIN (SELECT * FROM teams WHERE (user_one_id = @user1 OR user_two_id = @user1) AND (user_one_id = @user2 OR user_two_id = @user2) AND score = @score) as team1 ON matches.team_one_id = team1.id OR matches.team_two_id = team1.id", sql.Named("user1", playerOneID), sql.Named("user2", playerTwoID), sql.Named("score", teamOneScore)).
		Joins("JOIN (SELECT * FROM teams WHERE (user_one_id = @user3 OR user_two_id = @user3) AND (user_one_id = @user4 OR user_two_id = @user4) AND score = @score) as team2 ON matches.team_one_id = team2.id OR matches.team_two_id = team2.id", sql.Named("user3", playerThreeID), sql.Named("user4", playerFourID), sql.Named("score", teamTwoScore)).
		Where("matches.created_at > NOW() - INTERVAL '10 minutes'").
		Count(&count)

	return count > 0
}
