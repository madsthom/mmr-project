import { apiClient } from '$lib/server/api/apiClient';
import { fail, redirect } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import type { Actions, PageServerLoad } from './$types';
import { matchSchema } from './match-schema';

export const load: PageServerLoad = async ({ url }) => {
  const player1 = url.searchParams.get('player1') ?? '';
  const player2 = url.searchParams.get('player2') ?? '';
  const player3 = url.searchParams.get('player3') ?? '';
  const player4 = url.searchParams.get('player4') ?? '';

  return {
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
    const form = await superValidate(event, zod(matchSchema));

    if (!form.valid) {
      return fail(400, {
        form,
      });
    }

    try {
      await apiClient.mmrApi.mmrMatchesPost({ match: form.data });
    } catch (error) {
      return fail(500, {
        error,
      });
    }

    throw redirect(303, '/');
  },
};
