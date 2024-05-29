import { fail } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals: { apiClient } }) => {
  try {
    const entries = await apiClient.leaderboardApi.statsLeaderboardGet();
    const matches = await apiClient.mmrApi.mmrMatchesGet({
      limit: 5,
      offset: 0,
    });

    return {
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
