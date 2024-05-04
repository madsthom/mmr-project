import { apiClient } from '$lib/server/api/apiClient';
import { fail } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import type { Actions, PageServerLoad } from './$types';
import { matchSchema } from './match-schema';

export const load: PageServerLoad = async () => {
  return {
    form: await superValidate(zod(matchSchema)),
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
      await apiClient.mmrApi.mmrMatchPost({ match: form.data });
      return { form };
    } catch (error) {
      return fail(500, {
        error,
      });
    }
  },
};
