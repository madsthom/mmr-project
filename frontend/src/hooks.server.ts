import { createServerClient } from '@supabase/ssr';
import { redirect, type Handle } from '@sveltejs/kit';
import { sequence } from '@sveltejs/kit/hooks';

import {
  SUPERTOKENS_API_BASE_PATH,
  SUPERTOKENS_API_DOMAIN,
  SUPERTOKENS_API_KEY,
  SUPERTOKENS_CONNECTION_URI,
  SUPERTOKENS_WEBSITE_DOMAIN,
} from '$env/static/private';
import {
  PUBLIC_SUPABASE_ANON_KEY,
  PUBLIC_SUPABASE_URL,
} from '$env/static/public';
import { createApiClient } from '$lib/server/api/apiClient';

import SuperTokens from 'supertokens-node';
import Passwordless from 'supertokens-node/recipe/passwordless';
import Session from 'supertokens-node/recipe/session';

SuperTokens.init({
  supertokens: {
    connectionURI: SUPERTOKENS_CONNECTION_URI,
    apiKey: SUPERTOKENS_API_KEY,
  },
  appInfo: {
    appName: 'MMR Project',
    websiteDomain: SUPERTOKENS_WEBSITE_DOMAIN,
    apiDomain: SUPERTOKENS_API_DOMAIN,
    apiBasePath: SUPERTOKENS_API_BASE_PATH,
  },
  recipeList: [
    Session.init(), // Initializes session features
    Passwordless.init({
      contactMethod: 'EMAIL',
      flowType: 'MAGIC_LINK',
    }), // Initializes passwordless signin
  ],
});

export const handle = (async ({ event, resolve }) => {
  try {
    const accessToken = event.cookies.get(authCookieNames.access) ?? '';
    const antiCsrfToken = event.cookies.get(authCookieNames.csrf);
    const session = await Session.getSessionWithoutRequestResponse(
      accessToken,
      antiCsrfToken
    );
    const userId = session.getUserId();

    event.locals.user = { id: userId };
    return resolve(event);
  } catch (error) {
    if (!Session.Error.isErrorFromSuperTokens(error)) {
      return new Response('An unexpected error occurred', { status: 500 });
    }

    const userNeedsSessionRefresh =
      error.type === Session.Error.TRY_REFRESH_TOKEN;

    const requestAllowed =
      publicPages.includes(
        event.url.pathname as (typeof publicPages)[number]
      ) ||
      (userNeedsSessionRefresh &&
        event.url.pathname === commonRoutes.refreshSession);

    if (requestAllowed) {
      event.locals.user = {};
      return resolve(event);
    }

    const { url } = event;
    const basePath = userNeedsSessionRefresh
      ? commonRoutes.refreshSession
      : commonRoutes.login;
    const returnUrl = encodeURI(`${url.pathname}${url.search}`);
    const redirectUrl = `${basePath}?returnUrl=${returnUrl}`;

    // Redirect the user to the proper auth page. Delete their tokens if they don't need to attempt a token refresh.
    const headers = userNeedsSessionRefresh
      ? new Headers()
      : createHeadersFromTokens({});
    headers.set('Location', redirectUrl);
    return new Response(null, {
      status: userNeedsSessionRefresh ? 307 : 303,
      headers,
    });
  }
}) satisfies Handle;

const apiCliented: Handle = async ({ event, resolve }) => {
  if (event.locals.session != null) {
    event.locals.apiClient = createApiClient(event.locals.session.access_token);
  }

  return resolve(event);
};

export const handle: Handle = sequence(supabase, authGuard, apiCliented);
