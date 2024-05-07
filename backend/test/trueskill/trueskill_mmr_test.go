package trueskill

import (
	"mmr/backend/mmr"
	"os"
	"testing"

	"github.com/mafredri/go-trueskill"
	"github.com/stretchr/testify/assert"
)

func TestTrueskillMMRCalculation(t *testing.T) {
	team1 := mmr.Team{
		Players: []mmr.Player{mmr.CreateNewPlayer("test1", trueskill.DefaultMu, trueskill.DefaultSigma), mmr.CreateNewPlayer("test2", trueskill.DefaultMu, trueskill.DefaultSigma)},
		Score:   10,
	}

	team2 := mmr.Team{
		Players: []mmr.Player{mmr.CreateNewPlayer("test3", trueskill.DefaultMu, trueskill.DefaultSigma), mmr.CreateNewPlayer("test4", trueskill.DefaultMu, trueskill.DefaultSigma)},
		Score:   7,
	}

	team1, team2 = mmr.CalculateNewMMR(&team1, &team2, false)

	assert.True(t, team1.Players[0].Player.Mu() > team2.Players[0].Player.Mu() && team1.Players[0].Player.Mu() > team2.Players[1].Player.Mu(), "team1 won")
	assert.True(t, team1.Players[1].Player.Mu() > team2.Players[0].Player.Mu() && team1.Players[1].Player.Mu() > team2.Players[1].Player.Mu(), "team1 won")
}

func TestMain(m *testing.M) {

	// Run tests
	exitCode := m.Run()

	// Exit with the appropriate exit code
	os.Exit(exitCode)
}
