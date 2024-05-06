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
	"github.com/mafredri/go-trueskill"
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

	user1 := getUser(json.Team1.Member1)
	player1 := mmr.CreateNewPlayer(user1.Name, float64(user1.Mu), user1.Sigma)
	user2 := getUser(json.Team1.Member2)
	player2 := mmr.CreateNewPlayer(user2.Name, float64(user2.Mu), user2.Sigma)

	team1 := mmr.Team{
		Players: []mmr.Player{player1, player2},
		Score:   int16(json.Team1.Score),
	}

	user3 := getUser(json.Team2.Member1)
	player3 := mmr.CreateNewPlayer(user3.Name, float64(user3.Mu), user3.Sigma)
	user4 := getUser(json.Team2.Member2)
	player4 := mmr.CreateNewPlayer(user4.Name, float64(user4.Mu), user4.Sigma)

	team2 := mmr.Team{
		Players: []mmr.Player{player3, player4},
		Score:   int16(json.Team2.Score),
	}

	ts := trueskill.New(trueskill.DrawProbabilityZero())

	team1, team2 = mmr.CalculateNewMMR(ts, &team1, &team2)

	user1.Mu = team1.Players[0].Player.Mu()
	user1.Sigma = team1.Players[0].Player.Sigma()
	user1.MMR = int(user1.Mu - 3*user1.Sigma)
	user2.Mu = team1.Players[1].Player.Mu()
	user2.Sigma = team1.Players[1].Player.Sigma()
	user2.MMR = int(user2.Mu - 3*user2.Sigma)
	user3.Mu = team2.Players[0].Player.Mu()
	user3.Sigma = team2.Players[0].Player.Sigma()
	user3.MMR = int(user3.Mu - 3*user3.Sigma)
	user4.Mu = team2.Players[1].Player.Mu()
	user4.Sigma = team2.Players[1].Player.Sigma()
	user4.MMR = int(user4.Mu - 3*user4.Sigma)

	tm1m1 := upsertUser(user1)
	tm1m2 := upsertUser(user2)
	tm2m1 := upsertUser(user3)
	tm2m2 := upsertUser(user4)

	dbteam1 := createTeam(tm1m1, tm1m2, uint(json.Team1.Score), json.Team1.Score > json.Team2.Score)
	dbteam2 := createTeam(tm2m1, tm2m2, uint(json.Team2.Score), json.Team2.Score > json.Team1.Score)

	fmt.Println(dbteam1, dbteam2)

	match := createMatch(dbteam1, dbteam2)

	fmt.Println(match)

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Match submitted: %v", json)})
}

// @Summary Get leaderboard stats
// @Description Get leaderboard stats including wins, loses, and MMR of users
// @Tags Leaderboard
// @Accept json
// @Produce json
// @Success 200 {array} repos.LeaderboardEntry
// @Router /stats/leaderboard [get]
func GetLeaderboard(c *gin.Context) {
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

func upsertUser(user *models.User) uint {
	userRepo := repos.NewUserRepository(database.DB)
	user, err := userRepo.SaveUser(user)
	if err != nil {
		panic("Failed to save user")
	}

	return user.ID
}

func getUser(userName string) *models.User {
	userRepo := repos.NewUserRepository(database.DB)
	user, err := userRepo.GetOrCreateByName(userName)
	if err != nil {
		panic("Failed to find user")
	}

	return user
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
			s.GET("/leaderboard", GetLeaderboard)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
