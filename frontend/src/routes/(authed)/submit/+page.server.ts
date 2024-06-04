import { fail, redirect } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import type { ViewUserDetails } from '../../../api';
import type { Actions, PageServerLoad } from './$types';
import { matchSchema } from './match-schema';

const playerId = (
  url: URL,
  idParam: string,
  nameParam: string,
  users: ViewUserDetails[]
): number | undefined => {
  const id = url.searchParams.get(idParam);
  if (id != null) {
    const idNumber = parseInt(id, 10);
    if (!isNaN(idNumber)) {
      return idNumber;
    }
  }

  const name = url.searchParams.get(nameParam);
  return name != null
    ? users.find((user) => user.name === name)?.userId
    : undefined;
};

export const load: PageServerLoad = async ({ locals: { apiClient }, url }) => {
  const users = await apiClient.usersApi.v1UsersGet(); // TODO: Add error handling

  const player1 = playerId(url, 'player1_id', 'player1', users);
  const player2 = playerId(url, 'player2_id', 'player2', users);
  const player3 = playerId(url, 'player3_id', 'player3', users);
  const player4 = playerId(url, 'player4_id', 'player4', users);

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

    try {
      await apiClient.mmrApi.v2MmrMatchesPost({ match: form.data });
    } catch (error) {
      return fail(500, {
        error,
      });
    }

    throw redirect(303, '/');
  },
};
