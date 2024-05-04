import { env } from '$env/dynamic/private';
import { Configuration, DefaultApi, LeaderboardApi } from '../../../api';

const configuration = new Configuration({
  basePath: env.API_BASE_PATH,
});

export const apiClient = {
  mmrApi: new DefaultApi(configuration),
  leaderboardApi: new LeaderboardApi(configuration),
};
