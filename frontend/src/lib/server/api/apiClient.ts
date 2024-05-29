import { env } from '$env/dynamic/private';
import { Configuration, DefaultApi, LeaderboardApi } from '../../../api';

export const createConfiguration = (token: string) =>
  new Configuration({
    basePath: env.API_BASE_PATH,
    headers: { Authorization: `Bearer ${token}` },
  });

export const createApiClient = (token: string) => {
  const configuration = createConfiguration(token);
  return {
    leaderboardApi: new LeaderboardApi(configuration),
    mmrApi: new DefaultApi(configuration),
  };
};
