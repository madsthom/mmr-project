<script lang="ts">
  import Kpi from '$lib/components/kpi.svelte';
  import { MatchCard } from '$lib/components/match-card';
  import PageTitle from '$lib/components/page-title.svelte';
  import LineChart from '$lib/components/ui/line-chart/line-chart.svelte';
  import type { PageData } from './$types';

  export let data: PageData;
</script>

{#if data.user?.displayName}
  <PageTitle>{data.user?.displayName} ({data.user?.name})</PageTitle>
{:else}
  <PageTitle>{data.user?.name}</PageTitle>
{/if}

<div class="mt-6 grid grid-cols-[repeat(auto-fill,minmax(100px,1fr))] gap-2">
  <Kpi title="MMR">{data.mmrHistory[data.mmrHistory.length - 1].mmr}</Kpi>
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
  <Kpi title="Last match">
    {new Intl.RelativeTimeFormat(undefined, {
      style: 'narrow',
      numeric: 'auto',
    }).format(data.stats.daysSinceLastMatch, 'day')}
  </Kpi>
</div>

<h2 class="-mb-6 mt-6 text-2xl md:text-4xl">Rating over time</h2>

<LineChart
  data={data.mmrHistory.map((stat) => ({
    date: stat.date,
    player: stat.name,
    rating: stat.mmr,
  }))}
  height={300}
  legend={false}
/>

<h2 class="my-6 text-2xl md:text-4xl">Matches</h2>
<div class="flex flex-1 flex-col items-stretch gap-2">
  {#each data.matches ?? [] as match}
    <MatchCard users={data.users ?? []} {match} showMmr />
  {/each}
</div>
