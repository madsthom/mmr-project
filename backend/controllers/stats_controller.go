package controllers

import (
	"fmt"
	database "mmr/backend/db"
	"mmr/backend/db/repos"
	view "mmr/backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StatsController struct{}

//	@BasePath	/api/v1

// GetLeaderboard godoc
//
//	@Summary		Get leaderboard stats
//	@Description	Get leaderboard stats including wins, loses, and MMR of users
//	@Tags 			Leaderboard
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	repos.LeaderboardEntry
//	@Router			/stats/leaderboard [get]
func (sc StatsController) GetLeaderboard(c *gin.Context) {
	// Initialize leaderboard repository
	leaderboardRepo := repos.NewLeaderboardRepository(database.DB)

	// Fetch leaderboard entries
	entries, err := leaderboardRepo.GetLeaderboard()
	fmt.Println(entries)
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
//	@Param			userId	path	int	true	"User ID"
//	@Param			start	query	string	false	"Start date"
//	@Param			end		query	string	false	"End date"
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	view.PlayerHistoryDetails
//	@Router			/stats/player-history/{userId} [get]
func (sc StatsController) GetPlayerHistory(c *gin.Context) {
	// Initialize user repository
	userRepo := repos.NewUserRepository(database.DB)

	// Parse user ID from request as uint
	// If no user ID is provided, default to nil which will fetch all user histories
	var userId *uint = nil
	userIdParam := c.Param("userId")

	if userIdParam != "" {
		userIdParsed, err := strconv.ParseUint(userIdParam, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}
		userIdUint := uint(userIdParsed) // Convert userIdParsed to uint
		userId = &userIdUint
	}

	// Fetch user history
	entries, err := userRepo.ListPlayerHistory(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch player history"})
		return
	}

	if len(entries) == 0 {
		c.JSON(http.StatusOK, []view.PlayerHistoryDetails{})
		return
	}

	// Create list of view.PlayerHistoryDetails objects
	var playerHistory []view.PlayerHistoryDetails
	for _, entry := range entries {
		playerHistory = append(playerHistory, view.PlayerHistoryDetailsViewFromModel(*entry))
	}

	// Return player history as JSON response
	c.JSON(http.StatusOK, playerHistory)
}
