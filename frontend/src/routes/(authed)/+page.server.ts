import { fail } from '@sveltejs/kit';
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
      leaderboardEntries: entries.toSorted(
        (a, b) => (b.mmr ?? 0) - (a.mmr ?? 0)
      ),
      recentMatches: matches ?? [],
    };
  } catch (error) {
    fail(500, {
      message: 'Failed to load leaderboard',
    });
  }
};
