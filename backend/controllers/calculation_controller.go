package controllers

import (
	"fmt"
	"mmr/backend/mmr"
	view "mmr/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/intinig/go-openskill/rating"
	"github.com/intinig/go-openskill/types"
)

type CalculationController struct{}

// SubmitMMRCalculation godoc
//
//	@Summary		Submit an MMR calculation request
//	@Description	Submit two teams' details for MMR calculation
//	@Tags 			MMR
//	@Accept			json
//	@Produce		json
//	@Param			request	body		view.MMRCalculationRequest	true	"MMR Calculation Request"
//	@Success		200		{object}	view.MMRCalculationResponse	"MMR calculation result"
//	@Router			/v2/mmr/calculate [post]
func (m CalculationController) SubmitMMRCalculation(c *gin.Context) {
	var req view.MMRCalculationRequest
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ensurePlayers(c, req)

	// Create players for Team 1
	player1 := m.createPlayer(req.Team1.Players[0])
	player2 := m.createPlayer(req.Team1.Players[1])

	team1 := mmr.TeamV2{
		Players: []mmr.PlayerV2{player1, player2},
		Score:   int16(req.Team1.Score),
	}

	// Create players for Team 2
	player3 := m.createPlayer(req.Team2.Players[0])
	player4 := m.createPlayer(req.Team2.Players[1])

	team2 := mmr.TeamV2{
		Players: []mmr.PlayerV2{player3, player4},
		Score:   int16(req.Team2.Score),
	}

	// Calculate new MMR
	team1, team2 = mmr.CalculateNewMMRV2(&team1, &team2)

	// Prepare response
	response := view.MMRCalculationResponse{
		Team1: m.createTeamResult(req.Team1.Score, team1),
		Team2: m.createTeamResult(req.Team2.Score, team2),
	}

	// Respond with the updated team data
	c.JSON(http.StatusOK, response)
}

// Checks if the player IDs from both teams are unique and that there are exactly 4 unique players.
// If any validation fails, it responds with an appropriate error message and aborts the request.
func ensurePlayers(c *gin.Context, req view.MMRCalculationRequest) {
	// Extract all player IDs into a single slice
	var playerIDs []int64

	// Add all player IDs from Team 1
	for _, player := range req.Team1.Players {
		playerIDs = append(playerIDs, player.Id)
	}

	// Add all player IDs from Team 2
	for _, player := range req.Team2.Players {
		playerIDs = append(playerIDs, player.Id)
	}

	// Check for duplicates using a map
	playerMap := make(map[int64]struct{})
	for _, id := range playerIDs {
		if _, exists := playerMap[id]; exists {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Player ID %d is duplicated", id)})
			return
		}
		playerMap[id] = struct{}{}
	}

	// Ensure there are exactly 4 unique players
	if len(playerMap) != 4 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "There must be exactly 4 unique players across both teams"})
		return
	}
}


// Creates a player instance from the given MMRCalculationPlayerRating
func (m CalculationController) createPlayer(playerRating view.MMRCalculationPlayerRating) mmr.PlayerV2 {
	var internalRating types.Rating

	// Check if Mu and Sigma are provided; use defaults if they are nil
	if playerRating.Mu != nil && playerRating.Sigma != nil {
		// Create Rating with provided Mu and Sigma
		internalRating = types.Rating{
			Mu:    *playerRating.Mu,
			Sigma: *playerRating.Sigma,
			Z:     3, // Or set Z to a specific value if necessary
		}
	} else {
		// Use the New function to get a Rating with default options
		internalRating = rating.New()
	}

	return mmr.PlayerV2{
		Id: playerRating.Id,
		Player:   internalRating,
	}
}

// createTeamResult constructs the MMRTeamResult from score and calculated team data
func (m CalculationController) createTeamResult(score int, team mmr.TeamV2) view.MMRTeamResult {
	playersResults := make([]view.PlayerMMRResult, len(team.Players))

	for i, player := range team.Players {
		// Directly use the Mu, Sigma values from the team players
		playersResults[i] = view.PlayerMMRResult{
			Id:    player.Id, // Using Initials as the unique identifier
			Mu:    player.Player.Mu,
			Sigma: player.Player.Sigma,
			MMR:   float64(mmr.RankingDisplayValue(player.Player.Mu, player.Player.Sigma)),
		}
	}

	return view.MMRTeamResult{
		Score:   score,
		Players: playersResults,
	}
}