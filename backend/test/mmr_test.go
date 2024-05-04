package main

import (
	"fmt"
	"mmr/backend/mmr"
	"os"
	"testing"

	"github.com/mafredri/go-trueskill"
)

func TestMMRCalculation(t *testing.T) {

	ts := trueskill.New(trueskill.DrawProbabilityZero())
	p1 := ts.NewPlayer()
	p2 := trueskill.NewPlayer(p1.Mu()+5, p1.Sigma())
	p3 := ts.NewPlayer()
	p4 := ts.NewPlayer()

	team1 := mmr.NewTeam([]trueskill.Player{p1, p2})
	team2 := mmr.NewTeam([]trueskill.Player{p3, p4})

	avgPlayer1 := trueskill.NewPlayer(team1.GetMu(), team1.GetSigma())
	avgPlayer2 := trueskill.NewPlayer(team2.GetMu(), team2.GetSigma())

	skills := []trueskill.Player{avgPlayer1, avgPlayer2}

	newSkills, _ := ts.AdjustSkills(skills, false)
	avgnew1 := newSkills[0]
	avgnew2 := newSkills[1]

	fmt.Println("team1 1:", avgPlayer1.Mu()-avgnew1.Mu())
	fmt.Println("team2 2:", avgPlayer2.Mu()-avgnew2.Mu())
}

func TestMain(m *testing.M) {

	// Run tests
	exitCode := m.Run()

	// Exit with the appropriate exit code
	os.Exit(exitCode)
}
