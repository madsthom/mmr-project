package main

import (
	"fmt"
	"net/http"

	"mmr/backend/db/models"
	"mmr/backend/db/repos"
	docs "mmr/backend/docs"
	mmr "mmr/backend/mmr"

	database "mmr/backend/db"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@BasePath	/api/v1

// @Summary Submit a match
// @Description Submit a match for MMR calculation
// @Accept json
// @Produce json
// @Param match body mmr.Match true "Match object"
// @Success 200 {string} string "match submitted"
// @Router /mmr/match [post]
func SubmitMatch(c *gin.Context) {
	var json mmr.Match
	err := c.BindJSON(&json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. Calculate MMR
	// 2. Update users with new MMR
	// 3. Create two teams
	// 4. Create match with scores

	tm1m1 := upsertUser(json.Team1.Member1, 1)
	tm1m2 := upsertUser(json.Team1.Member2, 1)
	tm2m1 := upsertUser(json.Team2.Member1, 1)
	tm2m2 := upsertUser(json.Team2.Member2, 1)

	fmt.Println(tm1m1, tm1m2, tm2m1, tm2m2)

	team1 := createTeam(tm1m1, tm1m2, uint(json.Team1.Score), json.Team1.Score > json.Team2.Score)
	team2 := createTeam(tm2m1, tm2m2, uint(json.Team2.Score), json.Team2.Score > json.Team1.Score)

	fmt.Println(team1, team2)

	match := createMatch(team1, team2)

	fmt.Println(match)

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Match submitted: %v", json)})
}

// @Summary Get leaderboard stats
// @Description Get leaderboard stats including wins, loses, and MMR of users
// @Tags Leaderboard
// @Accept json
// @Produce json
// @Success 200 {array} LeaderboardEntry
// @Failure 500 {object} ErrorResponse
// @Router /stats/leaderboard [get]
func GetLeaderboard(c *gin.Context) {
    // Initialize leaderboard repository
    leaderboardRepo := repos.NewLeaderboardRepository(database.DB)
    
    // Fetch leaderboard entries
    entries, err := leaderboardRepo.GetLeaderboard()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch leaderboard entries"})
        return
    }
    
    // Return leaderboard entries as JSON response
    c.JSON(http.StatusOK, entries)
}

func createMatch(teamOneId, teamTwoId uint) uint {
	matchRepo := repos.NewMatchRepository(database.DB)
	match, err := matchRepo.CreateMatch(&models.Match{TeamOneID: teamOneId, TeamTwoID: teamTwoId})
	if err != nil {
		panic("Failed to create match")
	}
	return match.ID
}

func createTeam(playerOneId, playerTwoId, score uint, winner bool) uint {
	teamRepo := repos.NewTeamRepository(database.DB)
	team, err := teamRepo.CreateTeam(playerOneId, playerTwoId, score, winner)
	if err != nil {
		panic("Failed to create team")
	}
	return team.ID
}

func upsertUser(userName string, mmr int) uint {
	userRepo := repos.NewUserRepository(database.DB)
	user, err := userRepo.GetOrCreateByName(userName)
	if err != nil {
		panic("Failed to find user")
	} else {
		user.MMR = mmr
		userRepo.SaveUser(user)
	}

	return user.ID
}

func main() {
	database.InitDatabase()
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/mmr")
		{
			eg.POST("/match", SubmitMatch)
		}
		s := v1.Group("/stats")
		{
			s.GET("/leaderboard")
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
