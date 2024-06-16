import type { ViewMatchTeamV2, ViewUserDetails } from '$api';

export const load = async ({ url, locals: { apiClient } }) => {
  const player1String = url.searchParams.get('player1');
  const player2String = url.searchParams.get('player2');
  const opponent1String = url.searchParams.get('opponent1');
  const opponent2String = url.searchParams.get('opponent2');

  const player1Id = player1String ? parseInt(player1String) : null;
  const player2Id = player2String ? parseInt(player2String) : null;
  const opponent1Id = opponent1String ? parseInt(opponent1String) : null;
  const opponent2Id = opponent2String ? parseInt(opponent2String) : null;

  if (player1Id == null || player2Id == null) {
    throw new Error('Invalid players');
  }

  const [player1Matches, users] = await Promise.all([
    apiClient.mmrApi.v2MmrMatchesGet({
      userId: player1Id,
      limit: 1000,
      offset: 0,
    }),
    apiClient.usersApi.v1UsersGet(),
  ]);

  let matches = player1Matches.filter((match) => {
    return (
      (isOnTeam(match.team1, player1Id) && isOnTeam(match.team1, player2Id)) ||
      (isOnTeam(match.team2, player1Id) && isOnTeam(match.team2, player2Id))
    );
  });

  const player1 = users.find((user) => user.userId === player1Id);
  const player2 = users.find((user) => user.userId === player2Id);

  if (player1 == null || player2 == null) {
    throw new Error('Invalid players');
  }

  let opponents: [ViewUserDetails, ViewUserDetails] | null;
  if (opponent1Id != null) {
    if (opponent2Id == null) {
      throw new Error('Invalid opponents');
    }

    matches = matches.filter((match) => {
      return (
        (isOnTeam(match.team1, opponent1Id) &&
          isOnTeam(match.team1, opponent1Id)) ||
        (isOnTeam(match.team2, opponent2Id) &&
          isOnTeam(match.team2, opponent2Id))
      );
    });

    const opponent1 = users.find((user) => user.userId === opponent1Id);
    const opponent2 = users.find((user) => user.userId === opponent2Id);

    if (opponent1 == null || opponent2 == null) {
      throw new Error('Invalid opponents');
    }
    opponents = [opponent1, opponent2];
  } else {
    if (opponent2Id != null) {
      throw new Error('Invalid opponents');
    }
    opponents = null;
  }

  const totalMatches = matches.length;
  const wins = matches.filter((match) => {
    const winnerTeam =
      match.team1.score > match.team2.score ? match.team1 : match.team2;
    return isOnTeam(winnerTeam, player1Id);
  }).length;
  const lost = totalMatches - wins;
  const winrate = totalMatches > 0 ? wins / totalMatches : 0;

  const msSinceLastMatch = matches[0]
    ? new Date(matches[0].date).getTime() - new Date().getTime()
    : null;
  const daysSinceLastMatch = msSinceLastMatch
    ? millisecondsToDays(msSinceLastMatch)
    : null;

  return {
    users,
    players: [player1, player2],
    opponents,
    matches,
    stats:
      matches.length > 0
        ? {
            totalMatches,
            wins,
            lost,
            winrate,
            daysSinceLastMatch,
          }
        : null,
  };
};

const isOnTeam = (team: ViewMatchTeamV2, playerId: number) => {
  return team.member1 === playerId || team.member2 === playerId;
};

const millisecondsToDays = (ms: number) => Math.round(ms / 1000 / 60 / 60 / 24);
