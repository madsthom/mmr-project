<script lang="ts">
  export let delta: number;
</script>

<span
  class="mx-1 text-xs"
  class:positive={delta >= 0}
  class:destructive={delta < 0}
>
  ({delta > 0 ? '+' : ''}{delta})</span
>
