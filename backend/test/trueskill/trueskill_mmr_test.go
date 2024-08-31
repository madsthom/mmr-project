package trueskill

import (
	"mmr/backend/mmr"
	"os"
	"testing"

	"github.com/intinig/go-openskill/rating"
	"github.com/stretchr/testify/assert"
)

func TestTrueskillMMRCalculation(t *testing.T) {
	defaultRating := rating.New()
	team1 := mmr.Team{
		Players: []mmr.Player{mmr.CreateNewPlayer("test1", 27, defaultRating.Sigma-1), mmr.CreateNewPlayer("test2", defaultRating.Mu, defaultRating.Sigma)},
		Score:   10,
	}

	team2 := mmr.Team{
		Players: []mmr.Player{mmr.CreateNewPlayer("test3", defaultRating.Mu, defaultRating.Sigma), mmr.CreateNewPlayer("test4", defaultRating.Mu, defaultRating.Sigma)},
		Score:   7,
	}

	// log the player Mus and Sigma values for debugging purposes
	for _, player := range team1.Players {
		t.Logf("Team 1, Player %s MMR: %f, Mu: %f, Sigma: %f", player.Initials, mmr.RankingDisplayValue(player.Player.Mu, player.Player.Sigma), player.Player.Mu, player.Player.Sigma)
	}
	for _, player := range team2.Players {
		t.Logf("Team 2, Player %s MMR: %f, Mu: %f, Sigma: %f", player.Initials, mmr.RankingDisplayValue(player.Player.Mu, player.Player.Sigma), player.Player.Mu, player.Player.Sigma)
	}

	team1, team2 = mmr.CalculateNewMMR(&team1, &team2)

	// log the player Mus and Sigma values for debugging purposes
	for _, player := range team1.Players {
		t.Logf("Team 1, Player %s MMR: %f, Mu: %f, Sigma: %f", player.Initials, mmr.RankingDisplayValue(player.Player.Mu, player.Player.Sigma), player.Player.Mu, player.Player.Sigma)
	}
	for _, player := range team2.Players {
		t.Logf("Team 2, Player %s MMR: %f, Mu: %f, Sigma: %f", player.Initials, mmr.RankingDisplayValue(player.Player.Mu, player.Player.Sigma), player.Player.Mu, player.Player.Sigma)
	}

	assert.True(t, team1.Players[0].Player.Mu > team2.Players[0].Player.Mu && team1.Players[0].Player.Mu > team2.Players[1].Player.Mu, "team1 won")
	assert.True(t, team1.Players[1].Player.Mu > team2.Players[0].Player.Mu && team1.Players[1].Player.Mu > team2.Players[1].Player.Mu, "team1 won")
}

func TestMain(m *testing.M) {

	// Run tests
	exitCode := m.Run()

	// Exit with the appropriate exit code
	os.Exit(exitCode)
}
