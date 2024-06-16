<script lang="ts">
  import Kpi from '$lib/components/kpi.svelte';
  import { MatchCard } from '$lib/components/match-card';
  import PageTitle from '$lib/components/page-title.svelte';
  import Swords from 'lucide-svelte/icons/swords';

  export let data;
</script>

<div class="flex flex-col space-y-6">
  <PageTitle>Versus</PageTitle>

  <div class="flex items-center space-x-3">
    <div class="flex-1">
      <h2>Team1</h2>
      <div class="border-primary flex flex-col rounded-md border p-3">
        <p>{data.players[0].displayName ?? data.players[0].name}</p>
        <p>{data.players[1].displayName ?? data.players[1].name}</p>
      </div>
    </div>
    {#if data.opponents != null}
      <Swords />
      <div class="flex-1">
        <h2>Team2</h2>
        <div class="border-primary flex flex-col rounded-md border p-3">
          <p>{data.opponents[0].displayName ?? data.opponents[0].name}</p>
          <p>{data.opponents[1].displayName ?? data.opponents[1].name}</p>
        </div>
      </div>
    {/if}
  </div>

  {#if data.stats != null}
    <div class="grid grid-cols-[repeat(auto-fill,minmax(100px,1fr))] gap-2">
      <Kpi title="# Matches">
        {data.stats.totalMatches}
      </Kpi>
      <Kpi title="# Wins">
        {data.stats.wins}
      </Kpi>
      <Kpi title="# Losses">
        {data.stats.lost}
      </Kpi>
      <Kpi title="Win rate">
        {new Intl.NumberFormat(undefined, {
          style: 'percent',
          maximumFractionDigits: 1,
        }).format(data.stats.winrate)}
      </Kpi>
      {#if data.stats.daysSinceLastMatch != null}
        <Kpi title="Last match">
          {new Intl.RelativeTimeFormat(undefined, {
            style: 'narrow',
            numeric: 'auto',
          }).format(data.stats.daysSinceLastMatch, 'day')}
        </Kpi>
      {/if}
    </div>
  {/if}

  <div class="flex flex-col gap-3">
    <h2 class="text-2xl md:text-4xl">Matches</h2>
    <div class="flex flex-1 flex-col items-stretch gap-2">
      {#if data.matches.length === 0}
        <p>No matches found</p>
      {/if}
      {#each data.matches as match (match.date)}
        <MatchCard users={data.users ?? []} {match} showMmr />
      {/each}
    </div>
  </div>
</div>
