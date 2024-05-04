import { apiClient } from '$lib/server/api/apiClient';
import { fail } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
  try {
    const res = await apiClient.leaderboardApi.statsLeaderboardGet();
    
    return {
      leaderboardEntries: res.toSorted((a, b) => (b.mmr ?? 0) - (a.mmr ?? 0)),
    };
  } catch (error) {
    return fail(500, {
      message: 'Failed to load leaderboard',
    });
  }


};

