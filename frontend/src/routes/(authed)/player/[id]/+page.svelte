<script lang="ts">
  import Kpi from '$lib/components/kpi.svelte';
  import { MatchCard } from '$lib/components/match-card';
  import PageTitle from '$lib/components/page-title.svelte';
  import * as Card from '$lib/components/ui/card';
  import LineChart from '$lib/components/ui/line-chart/line-chart.svelte';
  import * as Table from '$lib/components/ui/table';
  import type { PageData } from './$types';

  export let data: PageData;

  const winRateFormatter = new Intl.NumberFormat(undefined, {
    style: 'percent',
    maximumFractionDigits: 0,
  });
</script>

<div class="flex flex-col gap-6">
  {#if data.user?.displayName}
    <PageTitle>{data.user?.displayName} ({data.user?.name})</PageTitle>
  {:else}
    <PageTitle>{data.user?.name}</PageTitle>
  {/if}

  <div class="grid grid-cols-[repeat(auto-fill,minmax(100px,1fr))] gap-2">
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

  {#if data.stats.opponents.length > 0}
    <Card.Root>
      <Card.Content class="flex flex-col p-0 md:p-6">
        <h2 class="text-xl md:text-2xl">Most common opponents</h2>
        <Table.Root class="">
          <Table.Header>
            <Table.Row>
              <!-- <Table.Head class="w-[3ch]">#</Table.Head> -->
              <Table.Head class="">Player</Table.Head>
              <Table.Head class="">
                <span class="sm:hidden">W</span>
                <span class="hidden sm:inline">Wins</span>
              </Table.Head>
              <Table.Head class="">
                <span class="sm:hidden">L</span>
                <span class="hidden sm:inline">Losses</span>
              </Table.Head>
              <Table.Head class="text-right">Win %</Table.Head>
            </Table.Row>
          </Table.Header>
          <Table.Body>
            {#each data.stats.opponents as { playerId, losses, wins, total }}
              {@const playerUser = data.users?.find(
                (user) => user.userId == playerId
              )}
              <Table.Row>
                <!-- <Table.Cell class="max-w-[3ch] font-bold">{rank}</Table.Cell> -->
                <Table.Cell>
                  <div class="flex flex-col items-start">
                    {#if playerUser?.displayName != null}
                      <span class="hidden w-full truncate sm:block">
                        {playerUser?.displayName}
                      </span>
                    {/if}
                    <span class="block">{playerUser?.name}</span>
                  </div>
                </Table.Cell>
                <Table.Cell>
                  {wins}
                </Table.Cell>
                <Table.Cell>
                  {losses}
                </Table.Cell>
                <Table.Cell class="text-right">
                  {winRateFormatter.format(wins / total)}
                </Table.Cell>
              </Table.Row>
            {/each}
          </Table.Body>
        </Table.Root>
      </Card.Content>
    </Card.Root>
  {/if}
  {#if data.stats.teammates.length > 0}
    <Card.Root>
      <Card.Content class="flex flex-col p-0 md:p-6">
        <h2 class="text-xl md:text-2xl">Most common teammates</h2>
        <Table.Root class="p-0">
          <Table.Header>
            <Table.Row>
              <!-- <Table.Head class="w-[3ch]">#</Table.Head> -->
              <Table.Head class="">Player</Table.Head>
              <Table.Head class="">
                <span class="sm:hidden">W</span>
                <span class="hidden sm:inline">Wins</span>
              </Table.Head>
              <Table.Head class="">
                <span class="sm:hidden">L</span>
                <span class="hidden sm:inline">Losses</span>
              </Table.Head>
              <Table.Head class="text-right">Win %</Table.Head>
            </Table.Row>
          </Table.Header>
          <Table.Body>
            {#each data.stats.teammates as { playerId, losses, wins, total }}
              {@const playerUser = data.users?.find(
                (user) => user.userId == playerId
              )}
              <Table.Row>
                <!-- <Table.Cell class="max-w-[3ch] font-bold">{rank}</Table.Cell> -->
                <Table.Cell>
                  <div class="flex flex-col items-start">
                    {#if playerUser?.displayName != null}
                      <span class="hidden w-full truncate sm:block">
                        {playerUser?.displayName}
                      </span>
                    {/if}
                    <span class="block">{playerUser?.name}</span>
                  </div>
                </Table.Cell>
                <Table.Cell>
                  {wins}
                </Table.Cell>
                <Table.Cell>
                  {losses}
                </Table.Cell>
                <Table.Cell class="text-right">
                  {winRateFormatter.format(wins / total)}
                </Table.Cell>
              </Table.Row>
            {/each}
          </Table.Body>
        </Table.Root>
      </Card.Content>
    </Card.Root>
  {/if}

  {#if data.matches.length > 0}
    <h2 class="my-6 text-2xl md:text-4xl">Matches</h2>
    <div class="flex flex-1 flex-col items-stretch gap-2">
      {#each data.matches ?? [] as match}
        <MatchCard users={data.users ?? []} {match} showMmr />
      {/each}
    </div>
  {/if}
</div>
