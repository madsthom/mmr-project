<script lang="ts">
  import * as Card from '$lib/components/ui/card';
  import * as Table from '$lib/components/ui/table';
  import type { LeaderboardEntry } from '../../routes/(authed)/+page.server';
  export let data: LeaderboardEntry[];

  const SHOW_STREAK_THRESHOLD = 3;
</script>

<Card.Root class="">
  <Card.Content>
    <Table.Root class="">
      <Table.Header>
        <Table.Row>
          <Table.Head class="w-16">#</Table.Head>
          <Table.Head class="">Player</Table.Head>
          <Table.Head class="">Wins</Table.Head>
          <Table.Head class="">Losses</Table.Head>
          <Table.Head class="text-right">Score</Table.Head>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {#each data as { loses, name, wins, mmr, mostRecentMMRChange, winningStreak, losingStreak }, idx}
          <Table.Row>
            <Table.Cell class="w-16 font-bold">{idx + 1}</Table.Cell>
            <Table.Cell>{name}</Table.Cell>
            <Table.Cell>
              {wins}
              {#if winningStreak && winningStreak >= SHOW_STREAK_THRESHOLD}
                <span class="ml-2 text-xs" title="Winning streak"
                  >🔥 {winningStreak}</span
                >
              {/if}
            </Table.Cell>
            <Table.Cell
              >{loses}
              {#if losingStreak && losingStreak >= SHOW_STREAK_THRESHOLD}
                <span class="ml-2 text-xs" title="Losing streak"
                  >🌧️ {losingStreak}</span
                >
              {/if}
            </Table.Cell>
            <Table.Cell class="text-right">
              {mmr != 0 ? mmr : '🐣'}
              {#if mostRecentMMRChange != null && mmr != 0}
                <span
                  title="MMR change at last match"
                  class={'text-xs ' +
                    (mostRecentMMRChange > 0
                      ? 'text-green-500'
                      : 'text-red-500')}
                >
                  ({mostRecentMMRChange > 0 ? '+' : ''}{mostRecentMMRChange})
                </span>
              {/if}
            </Table.Cell>
          </Table.Row>
        {/each}
      </Table.Body>
    </Table.Root>
  </Card.Content>
</Card.Root>
