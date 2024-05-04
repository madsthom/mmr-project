<script lang="ts">
  import Leaderboard from '$lib/components/leaderboard.svelte';
  import type { PageData } from './$types';

  export let data: PageData;
</script>

<div class="flex flex-col gap-4">
  <h1 class="text-center text-6xl">Trifork MMR</h1>

  <h2 class="text-4xl">Leaderboard</h2>
  <Leaderboard data={data.leaderboardEntries} />
</div>
