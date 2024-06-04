import { fail, redirect } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import type { Actions, PageServerLoad } from './$types';
import { createUserSchema } from './schema';

export const load: PageServerLoad = async ({ url }) => {
  const name = url.searchParams.get('name') ?? '';

  return {
    form: await superValidate(zod(createUserSchema), { defaults: { name } }),
  };
};

export const actions: Actions = {
  default: async (event) => {
    const apiClient = event.locals.apiClient;
    const form = await superValidate(event, zod(createUserSchema));

    const redirectTo = event.url.searchParams.get('redirect_to');

    if (!form.valid) {
      return fail(400, {
        form,
      });
    }

    try {
      const user = await apiClient.usersApi.v1UsersPost({
        user: form.data,
      });

      throw redirect(303, redirectTo != null ? redirectTo + user.userId : '/');
    } catch (error) {
      return fail(500, {
        error,
      });
    }
  },
};
