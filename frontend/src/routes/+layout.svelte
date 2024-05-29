<script lang="ts">
  import '../app.pcss';

  import { goto, invalidate } from '$app/navigation';
  import { onMount } from 'svelte';

  export let data;
  $: ({ session, supabase } = data);

  onMount(() => {
    const { data } = supabase.auth.onAuthStateChange((event, newSession) => {
      if (!newSession) {
        /**
         * Queue this as a task so the navigation won't prevent the
         * triggering function from completing
         */
        setTimeout(() => {
          goto('/', { invalidateAll: true });
        });
      }
      if (newSession?.expires_at !== session?.expires_at) {
        invalidate('supabase:auth');
      }

      if (event === 'SIGNED_IN') {
        goto('/', { invalidateAll: true });
      }
    });

    return () => data.subscription.unsubscribe();
  });
</script>

<slot />
