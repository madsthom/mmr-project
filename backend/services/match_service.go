package services

import (
	"errors"
	database "mmr/backend/db"
	"mmr/backend/db/models"
	"mmr/backend/db/repos"

	"github.com/mafredri/go-trueskill"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MatchService struct{}

func (ms MatchService) CreateMatch(seasonID uint, teamOneId, teamTwoId uint) uint {
	matchRepo := repos.NewMatchRepository(database.DB)
	match, err := matchRepo.CreateMatch(&models.Match{TeamOneID: teamOneId, TeamTwoID: teamTwoId, SeasonID: seasonID})
	if err != nil {
		panic("Failed to create match")
	}
	return match.ID
}

func (ms MatchService) CreateTeam(playerOneId, playerTwoId, score uint, winner bool) uint {
	teamRepo := repos.NewTeamRepository(database.DB)
	team, err := teamRepo.CreateTeam(playerOneId, playerTwoId, score, winner)
	if err != nil {
		panic("Failed to create team")
	}
	return team.ID
}

func (ms MatchService) CreatePlayerHistory(matchID uint, userID uint, mu float64, sigma float64, mmr int) uint {
	userRepo := repos.NewUserRepository(database.DB)
	playerHistory, err := userRepo.StoreRanking(matchID, userID, mu, sigma, mmr)
	if err != nil {
		panic("Failed to store player history")
	}
	return playerHistory.ID
}

func (ms MatchService) CreateMatchMMRCalculation(matchID uint, player1Delta int, player2Delta int, player3Delta int, player4Delta int) uint {
	userRepo := repos.NewUserRepository(database.DB)
	mmrCalculation, err := userRepo.StoreMatchMMRCalculation(matchID, player1Delta, player2Delta, player3Delta, player4Delta)
	if err != nil {
		panic("Failed to store MMR calculation")
	}

	return mmrCalculation.ID
}

func (ms MatchService) UpsertUser(user *models.User) uint {
	userRepo := repos.NewUserRepository(database.DB)
	user, err := userRepo.SaveUser(user)
	if err != nil {
		panic("Failed to save user")
	}

	return user.ID
}

func (ms MatchService) GetUser(userName string) *models.User {
	userRepo := repos.NewUserRepository(database.DB)
	user, err := userRepo.GetOrCreateByName(userName)
	if err != nil {
		panic("Failed to find user")
	}

	return user
}

func (ms MatchService) GetUserByID(userID uint) *models.User {
	userRepo := repos.NewUserRepository(database.DB)
	user, err := userRepo.GetByID(userID)
	if err != nil {
		panic("Failed to find user")
	}

	return user
}

func (ms MatchService) GetPlayerMuAndSigma(seasonID uint, userID uint) (Mu float64, Sigma float64) {
	userRepo := repos.NewUserRepository(database.DB)
	playerHistory, err := userRepo.GetLatestPlayerHistory(seasonID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return trueskill.DefaultMu, 2
		}
		panic("Failed to get player history")
	}

	return playerHistory.Mu, playerHistory.Sigma
}

func (ms MatchService) GetMatches(seasonID uint, limit int, offset int, orderByCreatedAtDesc bool, includeMmrCalculations bool, userId *uint) []*models.Match {
	matchRepo := repos.NewMatchRepository(database.DB)
	matches, err := matchRepo.ListMatches(seasonID, limit, offset, &clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: orderByCreatedAtDesc}, includeMmrCalculations, userId)

	if err != nil {
		panic("Failed to get matches")
	}

	return matches
}

func (ms MatchService) ClearAllMMRHistory(seasonID uint) {
	userRepo := repos.NewUserRepository(database.DB)
	if userRepo.ClearPlayerHistories(seasonID) != nil {
		panic("Failed to clear player histories")
	}

	matchRepo := repos.NewMatchRepository(database.DB)
	if matchRepo.ClearMMRCalculations(seasonID) != nil {
		panic("Failed to clear MMR calculations")
	}
}

func (ms MatchService) CheckExistingMatch(playerOneID uint, playerTwoID uint, playerThreeID uint, playerFourID uint, teamOneScore int, teamTwoScore int) bool {
	matchRepo := repos.NewMatchRepository(database.DB)
	return matchRepo.CheckExistingMatch(playerOneID, playerTwoID, playerThreeID, playerFourID, teamOneScore, teamTwoScore)
}
