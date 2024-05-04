<script lang="ts">
  import type { PageServerData } from './$types';
  import MatchForm from './components/match-form.svelte';

  export let data: PageServerData;
</script>

<div class="flex flex-col gap-8">
  <h1 class="font-extrabol text-center text-6xl">Submit match</h1>
  <MatchForm data={data.form} />
</div>
