package controllers

import (
	database "mmr/backend/db"
	"mmr/backend/db/models"
	"mmr/backend/db/repos"
	view "mmr/backend/models"
	"mmr/backend/services"
	queryParser "mmr/backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatsController struct{}

// GetLeaderboard godoc
//
//	@Summary		Get leaderboard stats
//	@Description	Get leaderboard stats including wins, loses, and MMR of users
//	@Tags 			Leaderboard
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	repos.LeaderboardEntry
//	@Router			/v1/stats/leaderboard [get]
func (sc StatsController) GetLeaderboard(c *gin.Context) {
	// Initialize leaderboard repository
	leaderboardRepo := repos.NewLeaderboardRepository(database.DB)

	seasonService := new(services.SeasonService)
	seasonID := seasonService.CurrentSeasonID()

	// Fetch leaderboard entries
	entries, err := leaderboardRepo.GetLeaderboard(seasonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch leaderboard entries"})
		return
	}

	// Return leaderboard entries as JSON response
	c.JSON(http.StatusOK, entries)
}

// GetPlayerHistory godoc
//
//	@Summary		Get player history
//	@Description	Get player history including MMR and date
//	@Tags 			Statistics
//	@Param			userId	query	int	false	"User ID"
//	@Param			start	query	string	false	"Start date"
//	@Param			end		query	string	false	"End date"
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	view.PlayerHistoryDetails
//	@Router			/v1/stats/player-history [get]
func (sc StatsController) GetPlayerHistory(c *gin.Context) {
	// Initialize user repository
	userRepo := repos.NewUserRepository(database.DB)
	seasonService := new(services.SeasonService)

	// If no user ID is provided, default to nil which will fetch all user histories
	userId := queryParser.GetNullableUintQueryValue(c, "userId")

	seasonID := seasonService.CurrentSeasonID()

	// Fetch user history
	entries, err := userRepo.ListPlayerHistory(&seasonID, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch player history"})
		return
	}

	if len(entries) == 0 {
		c.JSON(http.StatusOK, []view.PlayerHistoryDetails{})
		return
	}

	matchesPerUserID := make(map[uint]int)
	for _, entry := range entries {
		matchesPerUserID[entry.UserID]++
	}

	// Filter entries to include only those with userIDs that have 10 or more occurrences
	var filteredEntries []*models.PlayerHistory
	for _, entry := range entries {
		if matchesPerUserID[entry.UserID] >= 10 {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	if len(filteredEntries) == 0 {
		c.JSON(http.StatusOK, []view.PlayerHistoryDetails{})
		return
	}

	// Create list of view.PlayerHistoryDetails objects
	var playerHistory []view.PlayerHistoryDetails
	for _, entry := range filteredEntries {
		playerHistory = append(playerHistory, view.PlayerHistoryDetailsViewFromModel(*entry))
	}

	// Return player history as JSON response
	c.JSON(http.StatusOK, playerHistory)
}

// GetTimeStatistics godoc
//
//	@Summary		Get match distribution over time
//	@Description	Get number of matches for each day of week and hour of day
//	@Tags 			Statistics
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	view.TimeStatisticsEntry
//	@Router			/v1/stats/time-distribution [get]
func (sc StatsController) GetTimeStatistics(c *gin.Context) {

	matchRepo := repos.NewMatchRepository(database.DB)

	entries, err := matchRepo.GetMatchTimeDistribution()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate match distribution over time"})
		return
	}

	c.JSON(http.StatusOK, entries)
}
