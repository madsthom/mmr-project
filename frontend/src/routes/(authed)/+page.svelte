<script lang="ts">
  import Leaderboard from '$lib/components/leaderboard.svelte';
  import MatchCard from '$lib/components/match-card.svelte';
  import PageTitle from '$lib/components/page-title.svelte';
  import type { PageData } from './$types';

  export let data: PageData;
  const { leaderboardEntries, recentMatches } = data;
  let showMmr = false;
</script>

<div class="flex flex-col gap-4">
  <PageTitle>Trifork Foosball</PageTitle>
  <h2 class="text-4xl">Leaderboard</h2>
  <Leaderboard data={leaderboardEntries ?? []} />

  <div class="flex">
    <h2 class="flex-1 text-4xl">Recent Matches</h2>
    <div class="self-center">
      <label
        >MMR:&nbsp;
        <input type="checkbox" class="rounded-sm" bind:checked={showMmr} />
      </label>
    </div>
  </div>
  <div class="flex flex-1 flex-col items-stretch gap-2">
    {#each recentMatches ?? [] as match}
      <MatchCard {match} {showMmr} />
    {/each}
  </div>
</div>
