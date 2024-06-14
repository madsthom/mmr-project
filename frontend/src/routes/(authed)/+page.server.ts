import type { LeaderboardEntry } from '$lib/components/leaderboard/leader-board-entry';
import { fail } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals: { apiClient } }) => {
  try {
    // Not awaited by design, since we don't need to wait for this to render the page
    const statisticsPromise = apiClient.statisticsApi.v1StatsPlayerHistoryGet();

    const [entries, matches, users] = await Promise.all([
      apiClient.leaderboardApi.v1StatsLeaderboardGet(),
      apiClient.mmrApi.v2MmrMatchesGet({
        limit: 5,
        offset: 0,
      }),
      apiClient.usersApi.v1UsersGet(),
    ]);

    const leaderboardEntries = entries
      .toSorted((a, b) => (b.mmr ?? 0) - (a.mmr ?? 0))
      .map<LeaderboardEntry>((entry, idx) => ({ ...entry, rank: idx + 1 }));

    return {
      users,
      statisticsPromise,
      leaderboardEntries,
      recentMatches: matches ?? [],
    };
  } catch (error) {
    fail(500, {
      message: 'Failed to load leaderboard',
    });
  }
};
