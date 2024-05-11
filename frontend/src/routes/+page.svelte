<script lang="ts">
  import Leaderboard from '$lib/components/leaderboard.svelte';
  import PageTitle from '$lib/components/page-title.svelte';
  import type { PageData } from './$types';

  export let data: PageData;
</script>

<div class="flex flex-col gap-4">
  <PageTitle>Trifork Foosball</PageTitle>

  <h2 class="text-4xl">Leaderboard</h2>
  <Leaderboard data={data.leaderboardEntries ?? []} />
</div>
