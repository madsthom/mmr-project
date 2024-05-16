import { apiClient } from '$lib/server/api/apiClient';
import { fail } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
  try {
    const entries = await apiClient.leaderboardApi.statsLeaderboardGet();
    const matches = await apiClient.mmrApi.mmrMatchesGet();

    return {
      leaderboardEntries: entries.toSorted(
        (a, b) => (b.mmr ?? 0) - (a.mmr ?? 0)
      ),
      recentMatches:
        matches
          .toSorted(
            (a, b) => new Date(b.date).getTime() - new Date(a.date).getTime()
          )
          .slice(0, 5) ?? [],
    };
  } catch (error) {
    fail(500, {
      message: 'Failed to load leaderboard',
    });
  }
};
