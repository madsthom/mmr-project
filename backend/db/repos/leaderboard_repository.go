package repos

import (
	"mmr/backend/db/models"

	"gorm.io/gorm"
)

type LeaderboardEntry struct {
	Name  string `json:"name"`
	MMR   int    `json:"mmr"`
	Wins  int    `json:"wins"`
	Loses int    `json:"loses"`
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
			Name:  user.Name,
			MMR:   int(mapTrueSkillToMMR(user.Mu, user.Sigma)),
			Wins:  int(teamCounts.Winning),
			Loses: int(teamCounts.Losing),
		}

		// Append the entry to the results slice
		results = append(results, entry)
	}

	return results, nil
}

func mapTrueSkillToMMR(mu float64, sigma float64) float64 {
	// Define constants for TrueSkill to MMR mapping
	trueSkillMin := 0.0  // Minimum possible TrueSkill value (adjust as needed)
	trueSkillMax := 50.0 // Maximum possible TrueSkill value (adjust as needed)
	mmrMin := 1000.0     // Minimum MMR value
	mmrMax := 2000.0     // Maximum MMR value

	// Calculate the conservative estimate (mu - 3 * sigma)
	trueSkillValue := mu - 3.0*sigma

	// Perform linear interpolation to map TrueSkill value to MMR
	mmr := mmrMin + ((trueSkillValue-trueSkillMin)/(trueSkillMax-trueSkillMin))*(mmrMax-mmrMin)

	// Round MMR to the nearest integer
	return mmr
}
