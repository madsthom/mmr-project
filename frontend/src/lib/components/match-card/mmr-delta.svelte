<script lang="ts">
  export let delta: number | undefined | null = null;
</script>

{#if delta != null}
  <span
    class="text-xs"
    class:positive={delta >= 0}
    class:destructive={delta < 0}
  >
    ({delta > 0 ? '+' : ''}{delta})</span
  >
{/if}
