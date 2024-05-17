<script lang="ts">
  import Leaderboard from '$lib/components/leaderboard.svelte';
  import MatchCard from '$lib/components/match-card.svelte';
  import PageTitle from '$lib/components/page-title.svelte';
  import type { PageData } from './$types';

  export let data: PageData;
  const { leaderboardEntries, recentMatches } = data;
</script>

<div class="flex flex-col gap-4">
  <PageTitle>Trifork Foosball</PageTitle>
  <h2 class="text-4xl">Leaderboard</h2>
  <Leaderboard data={data.leaderboardEntries ?? []} />

  <h2 class="text-4xl">Recent Matches</h2>
  <div class="flex flex-1 flex-col items-stretch gap-2">
    {#each recentMatches ?? [] as match}
      <MatchCard {match} />
    {/each}
  </div>
</div>
