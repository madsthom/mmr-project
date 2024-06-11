<script lang="ts">
  import * as Card from '$lib/components/ui/card';
  import * as Table from '$lib/components/ui/table';
  import type { ReposLeaderboardEntry } from '../../api';
  export let data: ReposLeaderboardEntry[];

  const SHOW_STREAK_THRESHOLD = 3;
</script>

<Card.Root>
  <Card.Content class="p-0 md:p-6">
    <Table.Root class="">
      <Table.Header>
        <Table.Row>
          <Table.Head class="w-16">#</Table.Head>
          <Table.Head class="">Player</Table.Head>
          <Table.Head class="">
            <span class="sm:hidden">W</span>
            <span class="hidden sm:inline">Wins</span>
          </Table.Head>
          <Table.Head class="">
            <span class="sm:hidden">L</span>
            <span class="hidden sm:inline">Losses</span>
          </Table.Head>
          <Table.Head class="text-right">Score</Table.Head>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {#each data as { loses, name, wins, mmr, winningStreak, losingStreak }, idx}
          <Table.Row>
            <Table.Cell class="w-16 font-bold">{idx + 1}</Table.Cell>
            <Table.Cell>{name}</Table.Cell>
            <Table.Cell>
              <div class="flex flex-row items-center gap-2">
                {wins}
                {#if winningStreak && winningStreak >= SHOW_STREAK_THRESHOLD}
                  <span class="text-xs" title="Winning streak">
                    🔥 <span class="hidden sm:inline">{winningStreak}</span>
                  </span>
                {/if}
              </div>
            </Table.Cell>
            <Table.Cell>
              <div class="flex flex-row items-center gap-2">
                {loses}
                {#if losingStreak && losingStreak >= SHOW_STREAK_THRESHOLD}
                  <span class="text-xs" title="Losing streak">
                    🌧️ <span class="hidden sm:inline">{losingStreak}</span>
                  </span>
                {/if}
              </div>
            </Table.Cell>
            <Table.Cell class="text-right">{mmr != 0 ? mmr : '🐣'}</Table.Cell>
          </Table.Row>
        {/each}
      </Table.Body>
    </Table.Root>
  </Card.Content>
</Card.Root>
