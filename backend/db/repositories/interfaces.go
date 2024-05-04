// interfaces.go
package repositories

import "example.com/m/v2/db/models"

type IUserRepository interface {
    ListUsers() ([]*models.User, error)
}

type ITeamRepository interface {
    CreateTeam(team *models.Team) error
    ListTeams() ([]*models.Team, error)
}

type IMatchRepository interface {
    CreateMatch(match *models.Match) error
    ListMatches() ([]*models.Match, error)
}
