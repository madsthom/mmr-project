<script lang="ts">
  import * as Card from '$lib/components/ui/card';
  import * as Table from '$lib/components/ui/table';
  import type { ReposLeaderboardEntry } from '../../api';

  export let data: ReposLeaderboardEntry[];
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
        {#each data as { loses, name, wins, mmr }, idx}
          <Table.Row>
            <Table.Cell class="w-16 font-bold">{idx + 1}</Table.Cell>
            <Table.Cell>{name}</Table.Cell>
            <Table.Cell>{wins}</Table.Cell>
            <Table.Cell>{loses}</Table.Cell>
            <Table.Cell class="text-right"
              >{mmr != 0 ? mmr : 'Not enough matches'}</Table.Cell
            >
          </Table.Row>
        {/each}
      </Table.Body>
    </Table.Root>
  </Card.Content>
</Card.Root>
