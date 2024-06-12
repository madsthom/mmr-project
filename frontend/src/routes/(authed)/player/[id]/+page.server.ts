import { redirect } from '@sveltejs/kit';

export async function load({ params, locals: { apiClient } }) {
  const userId = Number(params.id);
  if (Number.isNaN(userId)) {
    throw redirect(303, '');
  }
  const user = await apiClient.usersApi.v1UsersIdGet({ id: userId });
  return { user };
}
