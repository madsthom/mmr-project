<script lang="ts">
  import '../../app.pcss';

  import Navbar from './components/navbar.svelte';

  export let data;
</script>

<main class="mx-auto mb-16 max-w-screen-sm overflow-auto p-4">
  <slot></slot>
</main>
<Navbar />

<style lang="postcss">
  :global(body) {
    @apply min-h-screen;
  }
</style>
