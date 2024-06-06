import { fail } from '@sveltejs/kit';
import type { ReposLeaderboardEntry, ViewMatchDetailsV2 } from '../../api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals: { apiClient } }) => {
  try {
    const [entries, matches, users] = await Promise.all([
      apiClient.leaderboardApi.v1StatsLeaderboardGet(),
      apiClient.mmrApi.v2MmrMatchesGet({
        limit: 5,
        offset: 0,
      }),
      apiClient.usersApi.v1UsersGet(),
    ]);

    return {
      users,
      leaderboardEntries: entries
        .toSorted((a, b) => (b.mmr ?? 0) - (a.mmr ?? 0))
        .map((entry) => ({
          ...entry,
          mostRecentMMRChange: getMostRecentMMRChange(matches[0], entry),
        })),
      recentMatches: matches ?? [],
    };
  } catch (error) {
    fail(500, {
      message: 'Failed to load leaderboard',
    });
  }
};

export type LeaderboardEntry = ReposLeaderboardEntry & {
  mostRecentMMRChange: ReturnType<typeof getMostRecentMMRChange>;
};

const getMostRecentMMRChange = (
  match: ViewMatchDetailsV2,
  player: ReposLeaderboardEntry
): number | undefined => {
  if (!match?.mmrCalculations) {
    return undefined;
  }
  if (match.team1.member1 === player.userId) {
    return match.mmrCalculations?.team1.player1MMRDelta;
  }
  if (match.team1.member2 === player.userId) {
    return match.mmrCalculations?.team1.player2MMRDelta;
  }
  if (match.team2.member1 === player.userId) {
    return match.mmrCalculations?.team2.player1MMRDelta;
  }
  if (match.team2.member2 === player.userId) {
    return match.mmrCalculations?.team2.player2MMRDelta;
  }
  return undefined;
};
