package controllers

import (
	"fmt"
	"mmr/backend/db/models"
	"mmr/backend/mmr"
	view "mmr/backend/models"
	services "mmr/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MatchController struct{}

//	@BasePath	/api/v1

// SubmitMatch godoc
//
//	@Summary		Submit a match
//	@Description	Submit a match for MMR calculation
//	@Accept			json
//	@Produce		json
//	@Param			match	body		view.Match	true	"Match object"
//	@Success		200		{string}	string		"match submitted"
//	@Router			/mmr/matches [post]
func (m MatchController) SubmitMatch(c *gin.Context) {
	var json view.Match
	err := c.BindJSON(&json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	matchService := new(services.MatchService)

	user1 := matchService.GetUser(json.Team1.Member1)
	player1 := m.createPlayer(matchService, user1)
	user2 := matchService.GetUser(json.Team1.Member2)
	player2 := m.createPlayer(matchService, user2)
	team1Score := *json.Team1.Score
	team2Score := *json.Team2.Score

	team1 := mmr.Team{
		Players: []mmr.Player{player1, player2},
		Score:   int16(team1Score),
	}

	user3 := matchService.GetUser(json.Team2.Member1)
	player3 := m.createPlayer(matchService, user3)
	user4 := matchService.GetUser(json.Team2.Member2)
	player4 := m.createPlayer(matchService, user4)

	team2 := mmr.Team{
		Players: []mmr.Player{player3, player4},
		Score:   int16(team2Score),
	}

	team1, team2 = mmr.CalculateNewMMR(&team1, &team2, false)

	user1.Mu = team1.Players[0].Player.Mu()
	user1.Sigma = team1.Players[0].Player.Sigma()
	user1.MMR = int(mmr.MapTrueSkillToMMR(user1.Mu, user1.Sigma))
	user2.Mu = team1.Players[1].Player.Mu()
	user2.Sigma = team1.Players[1].Player.Sigma()
	user2.MMR = int(mmr.MapTrueSkillToMMR(user2.Mu, user2.Sigma))
	user3.Mu = team2.Players[0].Player.Mu()
	user3.Sigma = team2.Players[0].Player.Sigma()
	user3.MMR = int(mmr.MapTrueSkillToMMR(user3.Mu, user3.Sigma))
	user4.Mu = team2.Players[1].Player.Mu()
	user4.Sigma = team2.Players[1].Player.Sigma()
	user4.MMR = int(mmr.MapTrueSkillToMMR(user4.Mu, user4.Sigma))

	tm1m1 := matchService.UpsertUser(user1)
	tm1m2 := matchService.UpsertUser(user2)
	tm2m1 := matchService.UpsertUser(user3)
	tm2m2 := matchService.UpsertUser(user4)

	dbteam1 := matchService.CreateTeam(tm1m1, tm1m2, team1Score, team1Score > team2Score)
	dbteam2 := matchService.CreateTeam(tm2m1, tm2m2, team2Score, team2Score > team1Score)

	fmt.Println(dbteam1, dbteam2)

	match := matchService.CreateMatch(dbteam1, dbteam2)

	fmt.Println(match)

	matchService.CreatePlayerHistory(match, user1.ID, user1.Mu, user1.Sigma, user1.MMR)
	matchService.CreatePlayerHistory(match, user2.ID, user2.Mu, user2.Sigma, user2.MMR)
	matchService.CreatePlayerHistory(match, user3.ID, user3.Mu, user3.Sigma, user3.MMR)
	matchService.CreatePlayerHistory(match, user4.ID, user4.Mu, user4.Sigma, user4.MMR)

	user1OldMMR := int(mmr.MapTrueSkillToMMR(player1.Player.Mu(), player1.Player.Sigma()))
	user2OldMMR := int(mmr.MapTrueSkillToMMR(player2.Player.Mu(), player2.Player.Sigma()))
	user3OldMMR := int(mmr.MapTrueSkillToMMR(player3.Player.Mu(), player3.Player.Sigma()))
	user4OldMMR := int(mmr.MapTrueSkillToMMR(player4.Player.Mu(), player4.Player.Sigma()))
	matchService.CreateMatchMMRCalculation(match, user1.MMR-user1OldMMR, user2.MMR-user2OldMMR, user3.MMR-user3OldMMR, user4.MMR-user4OldMMR)

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Match submitted: %v", json)})
}

// GetMatches
//
//	@Summary		Get matches
//	@Description	Get all matches
//	@Produce		json
//	@Success		200	{object}	[]view.MatchDetails
//	@Router			/mmr/matches [get]
func (m MatchController) GetMatches(c *gin.Context) {
	matchService := new(services.MatchService)

	matchesResult := matchService.GetMatches()

	var matches []view.MatchDetails
	for _, value := range matchesResult {
		match := view.MatchDetailsViewFromModel(*value)
		matches = append(matches, match)
	}

	c.JSON(http.StatusOK, matches)
}

func (m MatchController) createPlayer(matchService *services.MatchService, user *models.User) mmr.Player {
	Mu, Sigma := matchService.GetPlayerMuAndSigma(user.ID)
	return mmr.CreateNewPlayer(user.Name, Mu, Sigma)
}
