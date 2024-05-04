package main

import (
	"fmt"
	"mmr/backend/mmr"
	team "mmr/backend/mmr"
	"os"
	"testing"

	"github.com/mafredri/go-trueskill"
)

func createNewPlayer(initials string) mmr.Player {
	return mmr.Player{
		Initials: initials,
		Player:   trueskill.NewPlayer(trueskill.DefaultMu, trueskill.DefaultSigma),
	}
}

func createTeam(matchTeam mmr.MatchTeam) *mmr.Team {
	p1 := createNewPlayer(matchTeam.Member1)
	p2 := createNewPlayer(matchTeam.Member2)

	return &mmr.Team{
		Players: []mmr.Player{p1, p2},
		Score:   int16(matchTeam.Score),
	}
}

func getAverageTeamPlayer(team *mmr.Team) trueskill.Player {
	averageMu := team.GetAverageMu()
	averageSigma := team.GetAverageSigma()

	return trueskill.NewPlayer(averageMu, averageSigma)
}

func calculateNewMMR(ts trueskill.Config, team1 *mmr.Team, team2 *mmr.Team) (mmr.Team, mmr.Team) {
	avgTeam1 := getAverageTeamPlayer(team1)
	avgTeam2 := getAverageTeamPlayer(team2)

	fmt.Printf("Team1:\n")
	fmt.Printf("- Player1: %s - mu: %v\n", team1.Players[0].Initials, team1.Players[0].Player.Mu())
	fmt.Printf("- Player2: %s - mu: %v\n\n", team1.Players[1].Initials, team1.Players[1].Player.Mu())
	fmt.Printf("Team2:\n")
	fmt.Printf("- Player1: %s - mu: %v\n", team2.Players[0].Initials, team2.Players[0].Player.Mu())
	fmt.Printf("- Player2: %s - mu: %v\n\n", team2.Players[1].Initials, team2.Players[1].Player.Mu())

	fmt.Printf("AvgTeam1: %s\n", avgTeam1)
	fmt.Printf("AvgTeam2: %s\n\n", avgTeam2)

	var winner trueskill.Player
	var loser trueskill.Player
	if team1.Score > team2.Score {
		winner = avgTeam1
		loser = avgTeam2
	} else {
		winner = avgTeam2
		loser = avgTeam1
	}

	teams := []trueskill.Player{winner, loser}
	teamsWithNewSkills, _ := ts.AdjustSkills(teams, false)

	newTeam1 := teamsWithNewSkills[0]
	newTeam2 := teamsWithNewSkills[1]

	diffTeam1 := avgTeam1.Mu() - newTeam1.Mu()
	diffTeam2 := avgTeam2.Mu() - newTeam2.Mu()

	fmt.Printf("Team1: %s\n", newTeam1)
	fmt.Printf("Team2: %s\n", newTeam2)
	fmt.Printf("Team1 - Diff mu: %f\n", diffTeam1)
	fmt.Printf("Team2 - Diff mu: %f\n", diffTeam2)
	fmt.Println()

	return mmr.Team{}, mmr.Team{}
}

func TestMMRCalculation(t *testing.T) {
	ts := trueskill.New(trueskill.DrawProbabilityZero())

	match := team.Match{
		Team1: team.MatchTeam{
			Score:   10,
			Member1: "nila",
			Member2: "msth",
		},
		Team2: team.MatchTeam{
			Score:   8,
			Member1: "mbla",
			Member2: "maan",
		},
	}

	team1 := createTeam(match.Team1)
	team2 := createTeam(match.Team2)

	calculateNewMMR(ts, team1, team2)
}

func TestMain(m *testing.M) {

	// Run tests
	exitCode := m.Run()

	// Exit with the appropriate exit code
	os.Exit(exitCode)
}
