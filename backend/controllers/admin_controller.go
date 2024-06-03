package controllers

import (
	"mmr/backend/db/models"
	"mmr/backend/mmr"
	services "mmr/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

//	@BasePath	/api/v1

// RecalculateMatches godoc
//
//	@Summary		Recalculate matches
//	@Description	Start recalculating matches
//	@Tags 			Admin
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"recalculation done"
//	@Router			/admin/recalculate [post]
func (m AdminController) RecalculateMatches(c *gin.Context) {
	matchService := new(services.MatchService)

	matchService.ClearAllMMRHistory()
	var currentOffset = 0
	var limit = 200
	for {
		matches := matchService.GetMatches(limit, currentOffset, false, false)

		for _, match := range matches {
			dbTeam1 := match.TeamOne
			dbTeam2 := match.TeamTwo

			user1 := &dbTeam1.UserOne
			user2 := &dbTeam1.UserTwo
			user3 := &dbTeam2.UserOne
			user4 := &dbTeam2.UserTwo

			player1 := m.createPlayer(matchService, user1)
			player2 := m.createPlayer(matchService, user2)
			team1Score := dbTeam1.Score

			team1 := mmr.Team{
				Players: []mmr.Player{player1, player2},
				Score:   int16(team1Score),
			}

			player3 := m.createPlayer(matchService, user3)
			player4 := m.createPlayer(matchService, user4)
			team2Score := dbTeam2.Score

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

			_ = matchService.UpsertUser(user1)
			_ = matchService.UpsertUser(user2)
			_ = matchService.UpsertUser(user3)
			_ = matchService.UpsertUser(user4)

			matchService.CreatePlayerHistory(match.ID, user1.ID, user1.Mu, user1.Sigma, user1.MMR)
			matchService.CreatePlayerHistory(match.ID, user2.ID, user2.Mu, user2.Sigma, user2.MMR)
			matchService.CreatePlayerHistory(match.ID, user3.ID, user3.Mu, user3.Sigma, user3.MMR)
			matchService.CreatePlayerHistory(match.ID, user4.ID, user4.Mu, user4.Sigma, user4.MMR)

			user1OldMMR := int(mmr.MapTrueSkillToMMR(player1.Player.Mu(), player1.Player.Sigma()))
			user2OldMMR := int(mmr.MapTrueSkillToMMR(player2.Player.Mu(), player2.Player.Sigma()))
			user3OldMMR := int(mmr.MapTrueSkillToMMR(player3.Player.Mu(), player3.Player.Sigma()))
			user4OldMMR := int(mmr.MapTrueSkillToMMR(player4.Player.Mu(), player4.Player.Sigma()))

			matchService.CreateMatchMMRCalculation(match.ID, user1.MMR-user1OldMMR, user2.MMR-user2OldMMR, user3.MMR-user3OldMMR, user4.MMR-user4OldMMR)
		}

		if len(matches) == limit {
			currentOffset += limit
		} else {
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recalculated matches"})
}

func (m AdminController) createPlayer(matchService *services.MatchService, user *models.User) mmr.Player {
	Mu, Sigma := matchService.GetPlayerMuAndSigma(user.ID)
	return mmr.CreateNewPlayer(user.Name, Mu, Sigma)
}
