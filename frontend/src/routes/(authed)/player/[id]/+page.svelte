<script lang="ts">
  import type { PageData } from './$types';

  let data: PageData;
</script>

<h1>Player</h1>
<!-- Wins / Losses / Win% -->
<!-- Recent matches -->
<!-- MMR over time -->
<div class="flex flex-1 flex-col items-stretch gap-2">
  {#each data.matches ?? [] as match}
    <MatchCard users={data.users ?? []} {match} showMmr />
  {/each}
</div>
