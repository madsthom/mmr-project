import type { ViewMatchTeamV2 } from '../../../../api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({
  params,
  locals: { apiClient },
}) => {
  const playerId = Number(params.id);
  if (Number.isNaN(playerId)) {
    throw new Error('Invalid player ID');
  }
  const [matches, users, mmrHistory] = await Promise.all([
    apiClient.mmrApi.v2MmrMatchesGet({
      userId: playerId,
      limit: 1000,
      offset: 0,
    }),
    apiClient.usersApi.v1UsersGet(),
    apiClient.statisticsApi.v1StatsPlayerHistoryGet({ userId: playerId }),
  ]);

  const user = users.find((user) => user.userId === playerId);
  if (!user) {
    throw new Error('Player not found');
  }

  const totalMatches = matches.length;
  const wins = matches.filter((match) => {
    const winnerTeam =
      match.team1.score > match.team2.score ? match.team1 : match.team2;
    return isOnTeam(winnerTeam, playerId);
  }).length;
  const lost = totalMatches - wins;
  const winrate = totalMatches > 0 ? wins / totalMatches : 0;

  const msSinceLastMatch =
    new Date(matches[0].date).getTime() - new Date().getTime();
  const daysSinceLastMatch = millisecondsToDays(msSinceLastMatch);

  return {
    playerId,
    matches,
    users,
    user,
    mmrHistory,
    stats: {
      totalMatches,
      wins,
      lost,
      winrate,
      daysSinceLastMatch,
    },
  };
};

const isOnTeam = (team: ViewMatchTeamV2, playerId: number) => {
  return team.member1 === playerId || team.member2 === playerId;
};

const millisecondsToDays = (ms: number) => Math.round(ms / 1000 / 60 / 60 / 24);
