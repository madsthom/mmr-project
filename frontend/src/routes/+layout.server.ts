import type { LayoutServerLoad } from './$types';

export const load = (async ({ locals: { session } }) => {
  return {
    session,
  };
}) satisfies LayoutServerLoad;
