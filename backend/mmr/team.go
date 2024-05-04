package mmr

import (
	"fmt"
)

// Team is a composition of players that play together. The skill of a team
// (µ and σ) is determined by the skills of the players that form the team.
type Team struct {
	Players []Player
	Score   int16
}

// // NewEmptyTeam creates a team without any players. To add players use
// // AddPlayers.
// func NewEmptyTeam() Team {
// 	return Team{
// 		players: make([]trueskill.Player, 0),
// 	}
// }

// // NewTeam creates a team from a slice of players.
// func NewTeam(p []trueskill.Player) (t Team) {
// 	t.Players = p
// 	return
// }

// Size returns the number of players in the team
func (t *Team) Size() int {
	return len(t.Players)
}

// // AddPlayer adds a single player to the team.
// func (t *Team) AddPlayer(p trueskill.Player) {
// 	t.Players = append(t.Players, p)
// }

// // AddPlayers adds players to the team.
// func (t *Team) AddPlayers(p []trueskill.Player) {
// 	t.Players = append(t.Players, p...)
// }

// // GetPlayers returns the players the team is composed of.
// func (t *Team) GetPlayers() (p []trueskill.Player) {
// 	return t.Players
// }

// GetAverageMu returns the average of all means of the team
func (t *Team) GetAverageMu() (sum float64) {
	for _, p := range t.Players {
		sum += p.Player.Mu()
	}
	return sum / float64(t.Size())
}

// GetVar returns the combined variance of the team
func (t *Team) GetVar() (sum float64) {
	for _, p := range t.Players {
		sum += p.Player.Sigma() * p.Player.Sigma()
	}
	return
}

// GetAverageSigma returns the average of all sigma of the team
func (t *Team) GetAverageSigma() (sigma float64) {
	for _, p := range t.Players {
		sigma += p.Player.Sigma()
	}
	return sigma / float64(t.Size())
}

func (t *Team) String() (s string) {
	s = "Team of " + fmt.Sprint(len(t.Players)) + " Players:"
	for _, p := range t.Players {
		s += "\t" + p.Player.String()
	}
	return
}
