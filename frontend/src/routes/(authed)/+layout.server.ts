import type { LayoutServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';

export const load = (async ({ locals: { safeGetSession } }) => {
  const { session, user } = await safeGetSession();

  if (!session) {
    throw redirect(303, '/login');
  }

  return {
    session,
    user,
  };
}) satisfies LayoutServerLoad;
