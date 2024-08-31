package mmr

import (
	"github.com/intinig/go-openskill/rating"
	"github.com/intinig/go-openskill/types"
)

func CreateNewPlayer(initials string, mu float64, sigma float64) Player {
	return Player{
		Initials: initials,
		Player:   rating.NewWithOptions(&types.OpenSkillOptions{Mu: &mu, Sigma: &sigma}),
	}
}

func CalculateNewMMR(team1 *Team, team2 *Team) (Team, Team) {

	ratingResults := rating.Rate([]types.Team{
		{team1.Players[0].Player, team1.Players[1].Player},
		{team2.Players[0].Player, team2.Players[1].Player},
	}, &types.OpenSkillOptions{
		Score: []int{int(team1.Score), int(team2.Score)}, // it uses these scores to determine the winner
	})

	team1.Players[0].Player = ratingResults[0][0]
	team1.Players[1].Player = ratingResults[0][1]
	team2.Players[0].Player = ratingResults[1][0]
	team2.Players[1].Player = ratingResults[1][1]

	return *team1, *team2
}
