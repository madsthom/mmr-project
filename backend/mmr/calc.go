package mmr

import (
	"fmt"
	"math"

	"github.com/mafredri/go-trueskill"
)

func CreateNewPlayer(initials string, mu float64, sigma float64) Player {
	return Player{
		Initials: initials,
		Player:   trueskill.NewPlayer(mu, sigma),
	}
}

func GetAverageTeamPlayer(team *Team) trueskill.Player {
	averageMu := team.GetAverageMu()
	averageSigma := team.GetAverageSigma()

	return trueskill.NewPlayer(averageMu, averageSigma)
}

func CalculateNewMMR(ts trueskill.Config, team1 *Team, team2 *Team) (Team, Team) {
	avgTeam1 := GetAverageTeamPlayer(team1)
	avgTeam2 := GetAverageTeamPlayer(team2)

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
	var winnerTeam *Team
	var loserTeam *Team
	var win string
	if team1.Score > team2.Score {
		winner = avgTeam1
		loser = avgTeam2
		winnerTeam = team1
		loserTeam = team2
		win = "1"
	} else {
		winner = avgTeam2
		loser = avgTeam1
		winnerTeam = team2
		loserTeam = team1
		win = "2"
	}

	teams := []trueskill.Player{winner, loser}
	teamsWithNewSkills, _ := ts.AdjustSkills(teams, false)

	winnerAverageTeam := teamsWithNewSkills[0]
	loserAverageTeam := teamsWithNewSkills[1]

	winnerDiff := math.Abs(winner.Mu() - winnerAverageTeam.Mu())
	loserDiff := math.Abs(loser.Mu() - loserAverageTeam.Mu())

	fmt.Printf("Winner: %s\n", winnerAverageTeam)
	fmt.Printf("Loser: %s\n", loserAverageTeam)
	fmt.Printf("Winner - Diff mu: %f\n", winnerDiff)
	fmt.Printf("Loser - Diff mu: %f\n", loserDiff)
	fmt.Println()

	winnerTeam.Players[0].Player = trueskill.NewPlayer(winnerTeam.Players[0].Player.Mu()+winnerDiff, winnerAverageTeam.Sigma())
	winnerTeam.Players[1].Player = trueskill.NewPlayer(winnerTeam.Players[1].Player.Mu()+winnerDiff, winnerAverageTeam.Sigma())

	loserTeam.Players[0].Player = trueskill.NewPlayer(loserTeam.Players[0].Player.Mu()-loserDiff, loserAverageTeam.Sigma())
	loserTeam.Players[1].Player = trueskill.NewPlayer(loserTeam.Players[1].Player.Mu()-loserDiff, loserAverageTeam.Sigma())

	if win == "1" {
		return *winnerTeam, *loserTeam
	} else {
		return *loserTeam, *winnerTeam
	}

}
