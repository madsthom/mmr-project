package controllers

import (
	"fmt"
	"mmr/backend/db/models"
	"mmr/backend/mmr"
	view "mmr/backend/models"
	services "mmr/backend/services"
	queryParser "mmr/backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MatchController struct{}

// SubmitMatch godoc
//
//	@Summary		Submit a match
//	@Description	Submit a match for MMR calculation
//	@Tags 			Matches
//	@Accept			json
//	@Produce		json
//	@Param			match	body		view.Match	true	"Match object"
//	@Success		200		{string}	string		"match submitted"
//	@Router			/v1/mmr/matches [post]
func (m MatchController) SubmitMatch(c *gin.Context) {
	var json view.Match
	err := c.ShouldBind(&json)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// A match should have 4 unique players across 2 teams
	players := make(map[string]bool)
	players[json.Team1.Member1] = true
	players[json.Team1.Member2] = true
	players[json.Team2.Member1] = true
	players[json.Team2.Member2] = true

	if len(players) < 4 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Each player must be unique"})
		return
	}

	matchService := new(services.MatchService)
	seasonService := new(services.SeasonService)

	seasonID := seasonService.CurrentSeasonID()

	user1 := matchService.GetUser(json.Team1.Member1)
	player1 := m.createPlayer(seasonID, matchService, user1)
	user2 := matchService.GetUser(json.Team1.Member2)
	player2 := m.createPlayer(seasonID, matchService, user2)
	team1Score := *json.Team1.Score
	team2Score := *json.Team2.Score

	team1 := mmr.Team{
		Players: []mmr.Player{player1, player2},
		Score:   int16(team1Score),
	}

	user3 := matchService.GetUser(json.Team2.Member1)
	player3 := m.createPlayer(seasonID, matchService, user3)
	user4 := matchService.GetUser(json.Team2.Member2)
	player4 := m.createPlayer(seasonID, matchService, user4)

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

	match := matchService.CreateMatch(seasonID, dbteam1, dbteam2)

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

// SubmitMatchV2 godoc
//
//	@Summary		Submit a match
//	@Description	Submit a match for MMR calculation
//	@Tags 			Matches
//	@Accept			json
//	@Produce		json
//	@Param			match	body		view.MatchV2	true	"Match object"
//	@Success		200		{string}	string		"match submitted"
//	@Router			/v2/mmr/matches [post]
func (m MatchController) SubmitMatchV2(c *gin.Context) {
	var json view.MatchV2
	err := c.ShouldBind(&json)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// A match should have 4 unique players across 2 teams
	players := make(map[uint]bool)
	players[json.Team1.Member1] = true
	players[json.Team1.Member2] = true
	players[json.Team2.Member1] = true
	players[json.Team2.Member2] = true

	if len(players) < 4 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Each player must be unique"})
		return
	}

	matchService := new(services.MatchService)
	seasonService := new(services.SeasonService)

	team1Score := *json.Team1.Score
	team2Score := *json.Team2.Score

	// Check if there is an existing match
	if matchService.CheckExistingMatch(json.Team1.Member1, json.Team1.Member2, json.Team2.Member1, json.Team2.Member2, int(team1Score), int(team2Score)) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Match already exists"})
		return
	}

	seasonID := seasonService.CurrentSeasonID()

	user1 := matchService.GetUserByID(json.Team1.Member1)
	player1 := m.createPlayer(seasonID, matchService, user1)
	user2 := matchService.GetUserByID(json.Team1.Member2)
	player2 := m.createPlayer(seasonID, matchService, user2)

	team1 := mmr.Team{
		Players: []mmr.Player{player1, player2},
		Score:   int16(team1Score),
	}

	user3 := matchService.GetUserByID(json.Team2.Member1)
	player3 := m.createPlayer(seasonID, matchService, user3)
	user4 := matchService.GetUserByID(json.Team2.Member2)
	player4 := m.createPlayer(seasonID, matchService, user4)

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

	match := matchService.CreateMatch(seasonID, dbteam1, dbteam2)

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

// GetMatches godoc
//
//	@Summary		Get matches
//	@Description	Get all matches
//	@Tags 			Matches
//	@Param			limit	query	int	false	"Limit"
//	@Param			offset	query	int	false	"Offset"
//	@Produce		json
//	@Success		200	{object}	[]view.MatchDetails
//	@Router			/v1/mmr/matches [get]
func (m MatchController) GetMatches(c *gin.Context) {
	matchService := new(services.MatchService)
	seasonService := new(services.SeasonService)

	limitString := c.DefaultQuery("limit", "100")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	offsetString := c.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seasonID := seasonService.CurrentSeasonID()

	matchesResult := matchService.GetMatches(seasonID, limit, offset, true, true, nil)

	if len(matchesResult) == 0 {
		c.JSON(http.StatusOK, []view.MatchDetails{})
		return
	}

	var matches []view.MatchDetails
	for _, value := range matchesResult {
		match := view.MatchDetailsViewFromModel(*value)
		matches = append(matches, match)
	}

	c.JSON(http.StatusOK, matches)
}

// GetMatchesV2 godoc
//
//	@Summary		Get matches
//	@Description	Get all matches
//	@Tags 			Matches
//	@Param			limit	query	int	false	"Limit"
//	@Param			offset	query	int	false	"Offset"
//	@Param			userId	query	int	false	"User ID"
//	@Produce		json
//	@Success		200	{object}	[]view.MatchDetailsV2
//	@Router			/v2/mmr/matches [get]
func (m MatchController) GetMatchesV2(c *gin.Context) {
	matchService := new(services.MatchService)
	seasonService := new(services.SeasonService)

	limitString := c.DefaultQuery("limit", "100")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	offsetString := c.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// If no user ID is provided, default to nil which will fetch matches for all users
	userId := queryParser.GetNullableUintQueryValue(c, "userId")

	seasonID := seasonService.CurrentSeasonID()

	matchesResult := matchService.GetMatches(seasonID, limit, offset, true, true, userId)

	if len(matchesResult) == 0 {
		c.JSON(http.StatusOK, []view.MatchDetailsV2{})
		return
	}

	var matches []view.MatchDetailsV2
	for _, value := range matchesResult {
		match := view.MatchDetailsV2ViewFromModel(*value)
		matches = append(matches, match)
	}

	c.JSON(http.StatusOK, matches)
}

func (m MatchController) createPlayer(seasonID uint, matchService *services.MatchService, user *models.User) mmr.Player {
	Mu, Sigma := matchService.GetPlayerMuAndSigma(seasonID, user.ID)
	return mmr.CreateNewPlayer(user.Name, Mu, Sigma)
}
