import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({
  params,
  locals: { apiClient },
}) => {
  const playerId = Number(params.id);
  if (Number.isNaN(playerId)) {
    throw new Error('Invalid player ID');
  }
  const [matches, users] = await Promise.all([
    apiClient.mmrApi.v2MmrMatchesGet({
      userId: Number(playerId),
      limit: 100,
      offset: 0,
    }),
    apiClient.usersApi.v1UsersGet(),
  ]);

  return {
    playerId,
    matches,
    users,
  };
};
