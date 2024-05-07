package controllers

import (
	"fmt"
	database "mmr/backend/db"
	"mmr/backend/db/repos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LeaderboardController struct{}

// @Summary Get leaderboard stats
// @Description Get leaderboard stats including wins, loses, and MMR of users
// @Tags Leaderboard
// @Accept json
// @Produce json
// @Success 200 {array} repos.LeaderboardEntry
// @Router /stats/leaderboard [get]
func (l LeaderboardController) GetLeaderboard(c *gin.Context) {
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
