import { fail, redirect } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import type { Actions, PageServerLoad } from './$types';
import { matchSchema } from './match-schema';

export const load: PageServerLoad = async ({ locals: { apiClient }, url }) => {
  const users = await apiClient.usersApi.v1UsersGet(); // TODO: Add error handling
  const player1Param = url.searchParams.get('player1') ?? '';
  const player2Param = url.searchParams.get('player2') ?? '';
  const player3Param = url.searchParams.get('player3') ?? '';
  const player4Param = url.searchParams.get('player4') ?? '';

  const player1 = users.find((user) => user.name === player1Param)?.userId;
  const player2 = users.find((user) => user.name === player2Param)?.userId;
  const player3 = users.find((user) => user.name === player3Param)?.userId;
  const player4 = users.find((user) => user.name === player4Param)?.userId;

  return {
    users,
    form: await superValidate(zod(matchSchema), {
      defaults: {
        team1: { member1: player1, member2: player2, score: -1 },
        team2: { member1: player3, member2: player4, score: -1 },
      },
    }),
  };
};

export const actions: Actions = {
  default: async (event) => {
    const apiClient = event.locals.apiClient;
    const form = await superValidate(event, zod(matchSchema));

    if (!form.valid) {
      return fail(400, {
        form,
      });
    }

    console.log(form.data);

    try {
      await apiClient.mmrApi.v2MmrMatchesPost({ match: form.data });
    } catch (error) {
      console.log(error);
      return fail(500, {
        error,
      });
    }

    throw redirect(303, '/');
  },
};
