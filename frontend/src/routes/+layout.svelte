<script lang="ts">
  import '../app.pcss';

  import { invalidate } from '$app/navigation';
  import { onMount } from 'svelte';

  export let data;
  $: ({ session, supabase } = data);

  onMount(() => {
    const { data } = supabase.auth.onAuthStateChange((event, newSession) => {
      if (newSession?.expires_at !== session?.expires_at) {
        invalidate('supabase:auth');
      }
    });

    return () => data.subscription.unsubscribe();
  });
</script>

<slot />
