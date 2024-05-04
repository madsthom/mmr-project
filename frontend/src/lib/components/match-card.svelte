<script lang="ts">
  import * as Card from '$lib/components/ui/card';
  import type { MatchForm } from '../../routes/submit/match-schema';

  export let match: MatchForm;
</script>

<Card.Root>
  <div class="flex flex-row items-center gap-4 p-2 px-4">
    <div
      class="flex flex-1 flex-row items-center gap-4"
      class:text-primary={match.team1.score === 10}
    >
      <p class="text-4xl">
        {match.team1.score === 0 ? '🥚' : match.team1.score}
      </p>
      <div class="flex flex-1 flex-col">
        <p>{match.team1.member1}</p>
        <p>{match.team1.member2}</p>
      </div>
    </div>
    <div class="flex flex-col items-center">vs.</div>
    <div
      class="flex flex-1 flex-row items-center gap-4"
      class:text-primary={match.team2.score === 10}
    >
      <div class="flex flex-1 flex-col items-end">
        <p>{match.team2.member1}</p>
        <p>{match.team2.member2}</p>
      </div>
      <p class=" text-4xl">
        {match.team2.score === 0 ? '🥚' : match.team2.score}
      </p>
    </div>
  </div>
</Card.Root>
