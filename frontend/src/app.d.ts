// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces

import type { createApiClient } from '$lib/server/api/apiClient';
import type { Session, SupabaseClient } from '@supabase/supabase-js';

type ApiClient = ReturnType<typeof createApiClient>;

declare global {
  namespace App {
    // interface Error {}
    interface Locals {
      supabase: SupabaseClient;
      safeGetSession: () => Promise<{
        session: Session | null;
      }>;
      session: Session | null;
      apiClient: ApiClient;
    }
    // interface PageData {}
    // interface PageState {}
    // interface Platform {}
  }
}

export {};
