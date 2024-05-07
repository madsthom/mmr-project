package services

import (
	database "mmr/backend/db"
	"mmr/backend/db/models"
	"mmr/backend/db/repos"
)

type MatchService struct{}

func (ms MatchService) CreateMatch(teamOneId, teamTwoId uint) uint {
	matchRepo := repos.NewMatchRepository(database.DB)
	match, err := matchRepo.CreateMatch(&models.Match{TeamOneID: teamOneId, TeamTwoID: teamTwoId})
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
