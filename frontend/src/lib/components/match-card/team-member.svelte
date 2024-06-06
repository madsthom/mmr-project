<script lang="ts">
  import type { ViewMatchDetailsV2 } from '../../../api';
  import type { MatchUser } from './match-user';
  import MmrDelta from './mmr-delta.svelte';

  export let match: Pick<
    ViewMatchDetailsV2,
    'mmrCalculations' | 'team1' | 'team2'
  >;
  export let users: MatchUser[];
  export let showMmr = false;
  export let team: 'team1' | 'team2';
  export let member: 'member1' | 'member2';

  let memberId = match[team][member];
  let memberName =
    users.find((user) => user.userId === memberId)?.name ?? 'Unknown';
  let delta = null;
  switch (member) {
    case 'member1':
      delta = match.mmrCalculations?.[team].player1MMRDelta;
      break;
    case 'member2':
      delta = match.mmrCalculations?.[team].player2MMRDelta;
      break;
    default:
      delta = null;
      break;
  }

  let align = team === 'team1' ? 'left' : 'right';
</script>

<p class="space-x-1">
  {#if showMmr && align === 'right'}<MmrDelta {delta} />{/if}<span
    >{memberName}</span
  >{#if showMmr && align === 'left'}<MmrDelta {delta} />{/if}
</p>
