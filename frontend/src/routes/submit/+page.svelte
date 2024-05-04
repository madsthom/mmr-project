<script lang="ts">
  import { superForm } from 'sveltekit-superforms';
  import { zodClient } from 'sveltekit-superforms/adapters';
  import type { PageServerData } from './$types';
  import { matchSchema } from './match-schema';
  import MatchForm from './components/match-form.svelte';

  export let data: PageServerData;
</script>

<main class="p-4">
  <div class="flex flex-col gap-8">
    <h1 class="font-extrabol text-center text-6xl">Submit match</h1>
    <MatchForm data={data.form} />
  </div>
</main>
