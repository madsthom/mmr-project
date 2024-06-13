import { redirect } from '@sveltejs/kit';

import { fail, superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import type { Actions, PageServerLoad } from './$types';
import { signupSchema } from './schema';

export const load: PageServerLoad = async () => {
  const form = await superValidate(zod(signupSchema));
  return { form };
};

export const actions: Actions = {
  default: async (event) => {
    const form = await superValidate(event, zod(signupSchema));

    if (!form.valid) {
      return fail(400, {
        form,
      });
    }

    const { error } = await event.locals.supabase.auth.signUp(form.data);
    if (error) {
      console.error(error);
      return fail(400, { form });
    } else {
      return redirect(303, '/');
    }
  },
};
