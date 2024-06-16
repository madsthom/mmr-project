import { fail } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals: { apiClient } }) => {
  try {
    const [statistics, timeDistribution] = await Promise.all([
      apiClient.statisticsApi
        .v1StatsPlayerHistoryGet()
        .then((res) => res.toSorted((a, b) => a.name.localeCompare(b.name))),
      apiClient.statisticsApi.v1StatsTimeDistributionGet(),
    ]);

    return {
      statistics,
      timeDistribution,
    };
  } catch (error) {
    fail(500, {
      message: 'Failed to load statistics',
    });
  }
};
