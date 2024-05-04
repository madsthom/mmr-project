package mmr

import (
	"fmt"

	"github.com/mafredri/go-trueskill"
)

// Team is a composition of players that play together. The skill of a team
// (µ and σ) is determined by the skills of the players that form the team.
type Team struct {
	players []trueskill.Player
}

// NewEmptyTeam creates a team without any players. To add players use
// AddPlayers.
func NewEmptyTeam() Team {
	return Team{
		players: make([]trueskill.Player, 0),
	}
}

// NewTeam creates a team from a slice of players.
func NewTeam(p []trueskill.Player) (t Team) {
	t.players = p
	return
}

// Size returns the number of players in the team
func (t *Team) Size() int {
	return len(t.players)
}

// AddPlayer adds a single player to the team.
func (t *Team) AddPlayer(p trueskill.Player) {
	t.players = append(t.players, p)
}

// AddPlayers adds players to the team.
func (t *Team) AddPlayers(p []trueskill.Player) {
	t.players = append(t.players, p...)
}

// GetPlayers returns the players the team is composed of.
func (t *Team) GetPlayers() (p []trueskill.Player) {
	return t.players
}

// GetMu returns the sum of all means of the team
func (t *Team) GetMu() (sum float64) {
	for _, p := range t.players {
		sum += p.Mu()
	}
	return sum / float64(t.Size())
}

// GetVar returns the combined variance of the team
func (t *Team) GetVar() (sum float64) {
	for _, p := range t.players {
		sum += p.Sigma() * p.Sigma()
	}
	return
}

func (t *Team) GetSigma() (sigma float64) {
	for _, p := range t.players {
		sigma += p.Sigma()
	}
	return sigma / float64(t.Size())
}

func (t *Team) String() (s string) {
	s = "Team of " + fmt.Sprint(len(t.GetPlayers())) + " Players:"
	for _, p := range t.GetPlayers() {
		s += "\t" + p.String()
	}
	return
}
