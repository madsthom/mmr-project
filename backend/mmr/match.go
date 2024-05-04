package mmr

type Match struct {
	Team1 MatchTeam `json:"team1" binding:"required"`
	Team2 MatchTeam `json:"team2" binding:"required"`
}

type MatchTeam struct {
	Score   uint   `json:"score"`
	Member1 string `json:"member1" binding:"required"`
	Member2 string `json:"member2" binding:"required"`
}

func NewMatch(team1 MatchTeam, team2 MatchTeam) Match {
	return Match{
		Team1: team1,
		Team2: team2,
	}
}
