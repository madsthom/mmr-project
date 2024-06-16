<script>
  import Kpi from '$lib/components/kpi.svelte';
  import PageTitle from '$lib/components/page-title.svelte';
  import Button from '$lib/components/ui/button/button.svelte';
  import * as Card from '$lib/components/ui/card';
  import { Info } from 'lucide-svelte';

  export let data;
</script>

<div class="flex flex-col gap-3">
  <PageTitle>Profile</PageTitle>

  {#if data.playerId == null}
    <Card.Root class="flex flex-col gap-3 border-white p-4">
      <div class="flex space-x-2 text-white">
        <Info /><span>You have not claimed a player yet</span>
      </div>
      <Button href="/profile/claim">Claim a player</Button>
    </Card.Root>
  {/if}

  {#if data.stats != null}
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
    <Button href={`/player/${data.playerId}`}>See more stats</Button>
  {/if}

  <form method="post" action="/signout" class="contents">
    <Button type="submit" variant="secondary">Logout</Button>
  </form>
</div>
