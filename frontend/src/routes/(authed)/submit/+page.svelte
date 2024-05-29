<script lang="ts">
  import PageTitle from '$lib/components/page-title.svelte';
  import type { PageServerData } from './$types';
  import MatchForm from './components/match-form.svelte';

  export let data: PageServerData;
</script>

<div class="flex flex-col gap-8">
  <PageTitle>Submit match</PageTitle>
  <MatchForm data={data.form} />
</div>
