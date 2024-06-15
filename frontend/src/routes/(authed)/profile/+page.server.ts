import type {
  ViewMatchDetailsV2,
  ViewMatchTeamV2,
  ViewPlayerHistoryDetails,
} from '$api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals: { apiClient } }) => {
  const { userId } = await apiClient.profileApi.v1ProfileGet();
  const [matches, mmrHistory] =
    userId != null
      ? await Promise.all([
          apiClient.mmrApi.v2MmrMatchesGet({
            userId,
            limit: 1000,
            offset: 0,
          }),
          apiClient.statisticsApi.v1StatsPlayerHistoryGet({ userId }),
        ])
      : [null, null];

  return {
    playerId: userId,
    matches,
    mmrHistory,
    stats:
      userId != null && matches != null && mmrHistory != null
        ? getStats(userId, matches, mmrHistory)
        : null,
  };
};

const getStats = (
  playerId: number,
  matches: ViewMatchDetailsV2[],
  mmrHistory: ViewPlayerHistoryDetails[]
) => {
  const totalMatches = matches?.length ?? 0;
  const wins =
    playerId == null
      ? 0
      : matches?.filter((match) => {
          const winnerTeam =
            match.team1.score > match.team2.score ? match.team1 : match.team2;
          return isOnTeam(winnerTeam, playerId);
        }).length ?? 0;
  const lost = totalMatches - wins;
  const winrate = totalMatches > 0 ? wins / totalMatches : 0;

  const mmr = mmrHistory?.[mmrHistory.length - 1]?.mmr ?? null;

  const msSinceLastMatch =
    matches != null && matches[0]
      ? new Date(matches[0].date).getTime() - new Date().getTime()
      : null;
  const daysSinceLastMatch = msSinceLastMatch
    ? millisecondsToDays(msSinceLastMatch)
    : null;

  return { totalMatches, wins, lost, winrate, mmr, daysSinceLastMatch };
};

const isOnTeam = (team: ViewMatchTeamV2, playerId: number) => {
  return team.member1 === playerId || team.member2 === playerId;
};

const millisecondsToDays = (ms: number) => Math.round(ms / 1000 / 60 / 60 / 24);
