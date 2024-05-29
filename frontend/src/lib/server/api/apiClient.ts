import { env } from '$env/dynamic/private';
import { Configuration, LeaderboardApi, MatchesApi } from '../../../api';

export const createConfiguration = (token: string) =>
  new Configuration({
    basePath: env.API_BASE_PATH,
    headers: { Authorization: `Bearer ${token}` },
  });

export const createApiClient = (token: string) => {
  const configuration = createConfiguration(token);
  return {
    leaderboardApi: new LeaderboardApi(configuration),
    mmrApi: new MatchesApi(configuration),
  };
};
