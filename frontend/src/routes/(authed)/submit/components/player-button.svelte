<script lang="ts">
  import type { ViewUserDetails } from '../../../../api';

  export let user: ViewUserDetails;
</script>

<button
  class="border-input hover:border-primary focus-visible:ring-ring ring-offset-background flex w-full flex-col gap-1 rounded-md border px-3 py-2 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2"
  type="button"
  on:click
>
  <p>{user.displayName ?? user.name}</p>
  <p class="text-xs">
    {#if user.displayName != null}
      {user.name}
    {/if}
  </p>
</button>
