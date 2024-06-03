<script lang="ts">
  import * as Card from '$lib/components/ui/card';
  import type { ViewMatchDetails } from '../../../api';
  import MmrDelta from './mmr-delta.svelte';
  import TeamMember from './team-member.svelte';

  export let match: Omit<ViewMatchDetails, 'date'> & {
    date: string | undefined;
  };

  export let showMmr = false;
</script>

<Card.Root>
  <div class="flex flex-row items-center gap-4 p-2 px-4">
    <div
      class="flex flex-1 flex-row items-center gap-4"
      class:text-primary={match.team1.score === 10}
    >
      <p class="min-w-[2ch] text-4xl font-extrabold">
        {match.team1.score === 0 ? '🥚' : match.team1.score}
      </p>
      <div class="flex flex-1 flex-col">
        <TeamMember {match} {showMmr} team="team1" member="member1" />
        <TeamMember {match} {showMmr} team="team1" member="member2" />
      </div>
    </div>
    <div class="flex flex-col items-center">
      vs.
      {#if match.date}
        <p class="text-muted-foreground">
          {new Date(match.date).toLocaleTimeString(undefined, {
            hour: '2-digit',
            minute: '2-digit',
          })}
        </p>
      {/if}
    </div>
    <div
      class="flex flex-1 flex-row items-center gap-4"
      class:text-primary={match.team2.score === 10}
    >
      <div class="flex flex-1 flex-col items-end">
        <TeamMember {match} {showMmr} team="team2" member="member1" />
        <TeamMember {match} {showMmr} team="team2" member="member2" />
      </div>
      <p class="min-w-[2ch] text-right text-4xl font-extrabold">
        {match.team2.score === 0 ? '🥚' : match.team2.score}
      </p>
    </div>
  </div>
</Card.Root>
