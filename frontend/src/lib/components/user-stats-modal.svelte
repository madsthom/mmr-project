<script lang="ts">
  import { Button } from '$lib/components/ui/button';
  import * as Dialog from '$lib/components/ui/dialog';
  import { SHOW_STREAK_THRESHOLD } from '$lib/constants';
  import type { ReposLeaderboardEntry, ViewUserDetails } from '../../api';

  export let user: ViewUserDetails;
  export let leaderboardEntry: ReposLeaderboardEntry | null | undefined;
  export let open: boolean;
  export let onOpenChange: (open: boolean) => void;

  const percentFormatter = new Intl.NumberFormat(undefined, {
    style: 'percent',
    minimumFractionDigits: 2,
  });

  const isOnWinningStreak =
    (leaderboardEntry?.winningStreak ?? 0) >= SHOW_STREAK_THRESHOLD;
  const isOnLosingStreak =
    (leaderboardEntry?.losingStreak ?? 0) >= SHOW_STREAK_THRESHOLD;
</script>

<Dialog.Root {open} {onOpenChange}>
  <Dialog.Content>
    <Dialog.Title class="flex gap-2">
      <span>{user.displayName ?? user.name}</span>
      <span>
        {#if isOnWinningStreak}🔥{/if}{#if isOnLosingStreak}🌧️{/if}
      </span>
    </Dialog.Title>
    <Dialog.Description class="flex flex-col gap-4">
      {#if leaderboardEntry != null}
        {@const totalGamesPlayed =
          (leaderboardEntry.wins ?? 0) + (leaderboardEntry.loses ?? 0)}
        <div class="grid grid-cols-[auto_minmax(0,_1fr)] gap-x-5 gap-y-1">
          <p>Wins:</p>
          <p>
            {leaderboardEntry.wins}
            {#if leaderboardEntry.winningStreak != null && isOnWinningStreak}
              - streak: {leaderboardEntry.winningStreak}
            {/if}
          </p>
          <p>Losses:</p>
          <p>
            {leaderboardEntry.loses}
            {#if leaderboardEntry.losingStreak != null && isOnLosingStreak}
              - streak: {leaderboardEntry.losingStreak}
            {/if}
          </p>
          <p>Total games played:</p>
          <p>{totalGamesPlayed}</p>
          <p>Win %:</p>
          <p>
            {percentFormatter.format(
              totalGamesPlayed > 0
                ? (leaderboardEntry.wins ?? 0) / totalGamesPlayed
                : 0
            )}
          </p>
          <p>MMR:</p>
          <p>{leaderboardEntry.mmr}</p>
          <!-- Latest match -->
        </div>
      {/if}
      <Button href={`/player/${user.userId}`}>More details</Button>
    </Dialog.Description>
  </Dialog.Content>
</Dialog.Root>
