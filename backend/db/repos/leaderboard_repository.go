package repos

import (
	"mmr/backend/db/models"

	"gorm.io/gorm"
)

type LeaderboardEntry struct {
    Name   string
    MMR    int
    Wins   int
    Loses  int
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
            Winning int64
            Losing  int64
        }

        // Count the number of winning and losing teams where the user is either UserOne or UserTwo
        if err := lr.db.Model(&models.Team{}).
            Select("SUM(CASE WHEN winner = true THEN 1 ELSE 0 END) AS winning, SUM(CASE WHEN winner = false THEN 1 ELSE 0 END) AS losing").
            Where("(user_one_id = ? OR user_two_id = ?)", user.ID, user.ID).
            Scan(&teamCounts).Error; err != nil {
            return nil, err
        }

        // Create a new LeaderboardEntry object and populate it with the user's information and the counts of wins and losses
        entry := &LeaderboardEntry{
            Name:   user.Name,
            MMR:    user.MMR,
            Wins:   int(teamCounts.Winning),
            Loses:  int(teamCounts.Losing),
        }

        // Append the entry to the results slice
        results = append(results, entry)
    }

    return results, nil
}