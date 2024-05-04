package mmr

type Match struct {
	Team1 MatchTeam `json:"team1"`
	Team2 MatchTeam `json:"team2"`
}

type MatchTeam struct {
	Score   int    `json:"score"`
	Member1 string `json:"member1"`
	Member2 string `json:"member2"`
}

func NewMatch(team1 MatchTeam, team2 MatchTeam) Match {
	return Match{
		Team1: team1,
		Team2: team2,
	}
}
