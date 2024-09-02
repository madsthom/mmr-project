package repos

import (
	"database/sql"
	"mmr/backend/db/models"
	view "mmr/backend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IMatchRepository interface {
	CreateMatch(match *models.Match) (*models.Match, error)
	ListMatches(seasonID uint, limit int, offset int, orderBy *clause.OrderByColumn, includeMmrCalculations bool, userId *uint) ([]*models.Match, error)
	ClearMMRCalculations(seasonID uint) error
	CheckExistingMatch(playerOneID uint, playerTwoID uint, playerThreeID uint, playerFourID uint, teamOneScore int, teamTwoScore int) bool
	GetMatchTimeDistribution() ([]*view.TimeStatisticsEntry, error)
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

func (mr *MatchRepository) ListMatches(seasonID uint, limit int, offset int, orderBy *clause.OrderByColumn, includeMmrCalculations bool, userId *uint) ([]*models.Match, error) {
	var matches []*models.Match

	if orderBy == nil {
		orderBy = &clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: false}
	}

	query := mr.db.Model(&models.Match{}).
		Preload("TeamOne.UserOne").
		Preload("TeamOne.UserTwo").
		Preload("TeamTwo.UserOne").
		Preload("TeamTwo.UserTwo").
		Where("season_id = ?", seasonID)

	if includeMmrCalculations {
		query = query.Preload("MMRCalculations")
	}

	if userId != nil {
		query = query.
			Joins("JOIN teams AS team_one ON matches.team_one_id = team_one.id").
			Joins("JOIN teams AS team_two ON matches.team_two_id = team_two.id").
			Where("team_one.user_one_id = @user OR team_one.user_two_id = @user OR team_two.user_one_id = @user OR team_two.user_two_id = @user", sql.Named("user", *userId))
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

func (mr *MatchRepository) ClearMMRCalculations(seasonID uint) error {
	// Remove all MMR calculations where the match belongs to seasonID
	err := mr.db.Where("match_id IN (SELECT id FROM matches WHERE season_id = ?)", seasonID).Delete(&models.MMRCalculation{}).Error
	return err
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

func (mr *MatchRepository) GetMatchTimeDistribution() ([]*view.TimeStatisticsEntry, error) {
	var timeStatistics []*view.TimeStatisticsEntry

	err := mr.db.Model(&models.Match{}).
		Select("EXTRACT(DOW FROM created_at) as day_of_week, EXTRACT(HOUR FROM created_at) as hour_of_day, COUNT(*) as count").
		Group("day_of_week, hour_of_day").
		Order("day_of_week, hour_of_day").
		Find(&timeStatistics).Error

	if err != nil {
		return nil, err
	}

	return timeStatistics, nil
}
