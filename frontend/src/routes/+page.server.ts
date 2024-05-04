import { apiClient } from '$lib/server/api/apiClient';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
  const res = await apiClient.leaderboardApi.statsLeaderboardGet();

  console.log('Leaderboard entries:', res);

  return {
    leaderboardEntries: res.toSorted((a, b) => (b.mmr ?? 0) - (a.mmr ?? 0)),
  };
};

