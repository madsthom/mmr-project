<script lang="ts">
  import Button from '$lib/components/ui/button/button.svelte';
  import { Input } from '$lib/components/ui/input';
  import { isPresent } from '$lib/util/isPresent';
  import X from 'lucide-svelte/icons/x';
  import type { ViewUserDetails } from '../../../../api';
  import PlayerButton from './player-button.svelte';

  export let label: string;
  export let userId: number | null;
  export let users: ViewUserDetails[];
  export let latestPlayerIds: number[];
  export let availableUsers: ViewUserDetails[] = [];

  const resetValue = () => {
    userId = null;
  };

  let filter = '';

  $: filteredUsers = users.filter((u) =>
    u.name.toLowerCase().includes(filter.toLowerCase())
  );
  $: user = users.find((u) => u.userId === userId);
  $: latestPlayers = latestPlayerIds
    .map((id) => availableUsers.find((u) => u.userId === id))
    .filter(isPresent)
    .slice(0, 4);

  const selectUser = (user: ViewUserDetails) => {
    userId = user.userId;
    filter = '';
  };
</script>

<div class="flex flex-col gap-2">
  {label}
  {#if userId == null}
    <Input bind:value={filter} placeholder="Filter..." autofocus />
    {#if filter.length === 0 && latestPlayers.length > 0}
      <p class="text-sm">Recent players</p>
      <ul>
        {#each latestPlayers as latestPlayer}
          <li class="mb-1 last:mb-0">
            <PlayerButton
              user={latestPlayer}
              on:click={() => selectUser(latestPlayer)}
            />
          </li>
        {/each}
      </ul>
    {/if}
    {#if filter.length > 1}
      <ul>
        {#if filteredUsers.length > 0}{#each filteredUsers as user}
            <li>
              <button
                class="border-input hover:border-primary focus-visible:ring-ring ring-offset-background flex w-full flex-col gap-1 rounded-md border px-3 py-2 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2"
                type="button"
                on:click={() => {
                  userId = user.userId;
                  filter = '';
                }}
              >
                <p>{user.displayName ?? user.name}</p>
                <p class="text-xs">
                  {#if user.displayName != null}
                    {user.name}
                  {/if}
                </p>
              </button>
            </li>
          {/each}
        {:else}
          <li>No users found</li>
        {/if}
      </ul>
    {/if}
  {:else}
    <div
      class="border-input flex w-full items-center gap-1 rounded-md border px-3 py-2"
    >
      <div class="flex flex-1 flex-col gap-2">
        {#if user != null}
          <p>
            {user.displayName ?? user.name}
          </p>
          <p class="text-xs">
            {#if user.displayName != null}
              {user.name}
            {:else}
              &nbsp;
            {/if}
          </p>
        {:else}
          <p>Unknown</p>
        {/if}
      </div>
      <Button class="h-6 w-6 rounded p-1 text-sm" on:click={resetValue}>
        <X class="h-5 w-5" />
      </Button>
    </div>
  {/if}
</div>
