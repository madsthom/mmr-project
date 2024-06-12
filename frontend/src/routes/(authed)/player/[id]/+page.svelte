<script lang="ts">
  import PageTitle from '$lib/components/page-title.svelte';

  export let data;
</script>

<PageTitle>{data.user.displayName ?? data.user.name}</PageTitle>
