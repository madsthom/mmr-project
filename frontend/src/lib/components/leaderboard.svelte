<script lang="ts">
  import * as Card from '$lib/components/ui/card';
  import * as Table from '$lib/components/ui/table';

  const data = [
    {
      id: 1,
      name: 'maan',
      score: 1500,
    },
    {
      id: 2,
      name: 'mbla',
      score: 1456,
    },
    {
      id: 3,
      name: 'msth',
      score: 1345,
    },
  ];
</script>

<Card.Root class="">
  <Card.Content>
    <Table.Root class="">
      <Table.Header>
        <Table.Row>
          <Table.Head class="w-16">#</Table.Head>
          <Table.Head class="">Player</Table.Head>
          <Table.Head class="text-right">Score</Table.Head>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {#each data as { id, name, score }}
          <Table.Row>
            <Table.Cell class="w-16 font-bold">{id}</Table.Cell>
            <Table.Cell>{name}</Table.Cell>
            <Table.Cell class="text-right">{score}</Table.Cell>
          </Table.Row>
        {/each}
      </Table.Body>
    </Table.Root>
  </Card.Content>
</Card.Root>
