// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces

import type { ApiClient } from '$lib/server/api/apiClient';
import type { Session, SupabaseClient } from '@supabase/supabase-js';

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
