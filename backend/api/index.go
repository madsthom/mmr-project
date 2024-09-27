package api

import (
	"encoding/json"
	"fmt"
	"mmr/backend/mmr"
	view "mmr/backend/models"
	"net/http"

	"github.com/intinig/go-openskill/rating"
	"github.com/intinig/go-openskill/types"
)

// Handler is the entry point for the Vercel serverless function
func Handler(w http.ResponseWriter, r *http.Request) {
	calculationController := CalculationController{}
	calculationController.SubmitMMRCalculation(w, r)
}

type CalculationController struct{}

func (m CalculationController) SubmitMMRCalculation(w http.ResponseWriter, r *http.Request) {
	var req view.MMRCalculationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ensure players are valid
	if err := ensurePlayers(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

	// Set the response header and encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ensurePlayers checks if the player IDs from both teams are unique and that there are exactly 4 unique players.
func ensurePlayers(req view.MMRCalculationRequest) error {
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
			return fmt.Errorf("Player ID %d is duplicated", id)
		}
		playerMap[id] = struct{}{}
	}

	// Ensure there are exactly 4 unique players
	if len(playerMap) != 4 {
		return fmt.Errorf("There must be exactly 4 unique players across both teams")
	}

	return nil
}

// createPlayer creates a player instance from the given MMRCalculationPlayerRating
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
		Id:      playerRating.Id,
		Player:  internalRating,
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


