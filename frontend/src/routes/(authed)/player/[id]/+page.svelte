<script lang="ts">
  import Kpi from '$lib/components/kpi.svelte';
  import { MatchCard } from '$lib/components/match-card';
  import PageTitle from '$lib/components/page-title.svelte';
  import * as Card from '$lib/components/ui/card';
  import LineChart from '$lib/components/ui/line-chart/line-chart.svelte';
  import { isPresent } from '$lib/util/isPresent';
  import { Separator } from 'bits-ui';
  import type { PageData } from './$types';

  export let data: PageData;

  $: teammates = data.matches
    .map((match) => {
      if (
        match.team1.member1 === data.user.userId ||
        match.team1.member2 === data.user.userId
      ) {
        return match.team1;
      }

      return match.team2;
    })
    .flatMap((team) => [team.member1, team.member2])
    .filter(isPresent)
    .reduce<Record<number, number>>((acc, player) => {
      if (player === data.user.userId) {
        return acc;
      }
      acc[player] = (acc[player] ?? 0) + 1;
      return acc;
    }, {});

  $: opponents = data.matches
    .map((match) => {
      if (
        match.team1.member1 !== data.user.userId &&
        match.team1.member2 !== data.user.userId
      ) {
        return match.team1;
      }

      return match.team2;
    })
    .flatMap((team) => [team.member1, team.member2])
    .filter(isPresent)
    .reduce<Record<number, number>>((acc, player) => {
      if (player === data.user.userId) {
        return acc;
      }
      acc[player] = (acc[player] ?? 0) + 1;
      return acc;
    }, {});

  $: mostPlayedOpponent = Object.entries(opponents).reduce<{
    player: string;
    count: number;
  } | null>((acc, [player, count]) => {
    if (acc == null || count > acc.count) {
      return { player, count };
    }

    return acc;
  }, null);

  $: mostPlayedTeammate = Object.entries(teammates).reduce<{
    player: string;
    count: number;
  } | null>((acc, [player, count]) => {
    if (acc == null || count > acc.count) {
      return { player, count };
    }

    return acc;
  }, null);
</script>

{#if data.user?.displayName}
  <PageTitle>{data.user?.displayName} ({data.user?.name})</PageTitle>
{:else}
  <PageTitle>{data.user?.name}</PageTitle>
{/if}

<div class="mt-6 grid grid-cols-[repeat(auto-fill,minmax(100px,1fr))] gap-2">
  {#if data.stats.mmr != null}
    <Kpi title="MMR">{data.stats.mmr}</Kpi>
  {/if}
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

<Card.Root class="mt-6 flex flex-col gap-4 bg-gray-800 p-4 shadow-sm">
  {#if mostPlayedOpponent != null}
    {@const mostPlayedOpponentUser = data.users?.find(
      (user) => user.userId.toString() === mostPlayedOpponent.player
    )}
    {#if mostPlayedOpponentUser != null}
      <div class="flex flex-col gap-2">
        <h3 class="text-lg text-gray-300">Biggest Enemy</h3>
        <p class="text-primary font-bold">
          {mostPlayedOpponentUser.displayName ?? mostPlayedOpponentUser.name} - {mostPlayedOpponent.count}
          matches
        </p>
      </div>
    {/if}
  {/if}
  <Separator.Root
    orientation="horizontal"
    class="bg-background shrink-0 data-[orientation=horizontal]:h-px data-[orientation=vertical]:h-full data-[orientation=horizontal]:w-full data-[orientation=vertical]:w-[1px]"
  />
  {#if mostPlayedTeammate != null}
    {@const mostPlayedTeammateUser = data.users?.find(
      (user) => user.userId.toString() === mostPlayedTeammate.player
    )}
    {#if mostPlayedTeammateUser != null}
      <div class="flex flex-col gap-2">
        <h3 class="text-lg text-gray-300">Best Friend</h3>
        <p class="text-primary font-bold">
          {mostPlayedTeammateUser.displayName ?? mostPlayedTeammateUser.name} - {mostPlayedTeammate.count}
          matches
        </p>
      </div>
    {/if}
  {/if}
</Card.Root>

{#if data.mmrHistory.length > 0}
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
{/if}

{#if data.matches.length > 0}
  <h2 class="my-6 text-2xl md:text-4xl">Matches</h2>
  <div class="flex flex-1 flex-col items-stretch gap-2">
    {#each data.matches ?? [] as match}
      <MatchCard users={data.users ?? []} {match} showMmr />
    {/each}
  </div>
{/if}
