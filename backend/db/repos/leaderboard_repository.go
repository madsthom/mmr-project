package repos

import (
	"mmr/backend/db/models"
	"mmr/backend/mmr"

	"gorm.io/gorm"
)

type LeaderboardEntry struct {
	UserID        uint   `json:"userId"`
	Name          string `json:"name"`
	MMR           int    `json:"mmr"`
	Wins          int    `json:"wins"`
	Loses         int    `json:"loses"`
	WinningStreak int    `json:"winningStreak"`
	LosingStreak  int    `json:"losingStreak"`
}

// ILeaderboardRepository interface declaration
type ILeaderboardRepository interface {
	// GetLeaderboard fetches users from the database along with the count of winning and losing teams they belong to
	GetLeaderboard() ([]*LeaderboardEntry, error)
}

// LeaderboardRepository struct
type LeaderboardRepository struct {
	db *gorm.DB
}

// NewLeaderboardRepository initializes a new LeaderboardRepository
func NewLeaderboardRepository(db *gorm.DB) ILeaderboardRepository {
	return &LeaderboardRepository{db}
}

// GetLeaderboard fetches users from the database along with the count of winning and losing teams they belong to
func (lr *LeaderboardRepository) GetLeaderboard() ([]*LeaderboardEntry, error) {
	var results []*LeaderboardEntry

	// Fetch users from the database
	var users []*models.User
	if err := lr.db.Find(&users).Error; err != nil {
		return nil, err
	}

	// Iterate over users and count the number of winning and losing teams they belong to
	for _, user := range users {
		var teamCounts struct {
			Winning       int64
			Losing        int64
			WinningStreak int64
			LosingStreak  int64
		}

		// Count the number of winning and losing teams where the user is either UserOne or UserTwo
		if err := lr.db.Model(&models.Team{}).
			Select("SUM(CASE WHEN winner = true THEN 1 ELSE 0 END) AS winning, SUM(CASE WHEN winner = false THEN 1 ELSE 0 END) AS losing").
			Where("(user_one_id = ? OR user_two_id = ?)", user.ID, user.ID).
			Scan(&teamCounts).Error; err != nil {
			return nil, err
		}

		// Count number of rows created after the last time the user lost a game
		lr.db.Model(&models.Team{}).
			Where("(user_one_id = ? OR user_two_id = ?)", user.ID, user.ID).
			Where("id > (SELECT MAX(id) from teams WHERE winner = false AND (user_one_id = ? OR user_two_id = ?))", user.ID, user.ID).
			Count(&teamCounts.WinningStreak)

		// If the user has never lost a game, set the winning streak to the number of wins
		if teamCounts.Losing == 0 {
			teamCounts.WinningStreak = teamCounts.Winning
		}

		// Count number of rows created after the last time the user won a game
		lr.db.Model(&models.Team{}).
			Where("(user_one_id = ? OR user_two_id = ?)", user.ID, user.ID).
			Where("id > (SELECT MAX(id) from teams WHERE winner = true AND (user_one_id = ? OR user_two_id = ?))", user.ID, user.ID).
			Count(&teamCounts.LosingStreak)

		// If the user has never won a game, set the losing streak to the number of losses
		if teamCounts.Winning == 0 {
			teamCounts.LosingStreak = teamCounts.Losing
		}

		// Create a new LeaderboardEntry object and populate it with the user's information and the counts of wins and losses
		entry := &LeaderboardEntry{
			UserID:        user.ID,
			Name:          user.Name,
			MMR:           int(mmr.RankingDisplayValue(user.Mu, user.Sigma)),
			Wins:          int(teamCounts.Winning),
			Loses:         int(teamCounts.Losing),
			WinningStreak: int(teamCounts.WinningStreak),
			LosingStreak:  int(teamCounts.LosingStreak),
		}

		if entry.Wins+entry.Loses < 10 {
			entry.MMR = 0
		}

		// Append the entry to the results slice
		results = append(results, entry)
	}

	return results, nil
}
