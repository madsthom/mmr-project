package view

import (
	"mmr/backend/db/models"
	"time"
)

type Match struct {
	Team1 MatchTeam `json:"team1" binding:"required"`
	Team2 MatchTeam `json:"team2" binding:"required"`
}

type MatchDetails struct {
	Date  time.Time `json:"date" binding:"required"`
	Team1 MatchTeam `json:"team1" binding:"required"`
	Team2 MatchTeam `json:"team2" binding:"required"`
}

type MatchTeam struct {
	Score   *uint  `json:"score" binding:"required"`
	Member1 string `json:"member1" binding:"required"`
	Member2 string `json:"member2" binding:"required"`
}

func MatchTeamViewFromModel(team models.Team) MatchTeam {
	return MatchTeam{
		Score:   &team.Score,
		Member1: team.UserOne.Name,
		Member2: team.UserTwo.Name,
	}
}

func MatchDetailsViewFromModel(match models.Match) MatchDetails {
	return MatchDetails{
		Team1: MatchTeamViewFromModel(match.TeamOne),
		Team2: MatchTeamViewFromModel(match.TeamTwo),
		Date:  match.CreatedAt,
	}
}
