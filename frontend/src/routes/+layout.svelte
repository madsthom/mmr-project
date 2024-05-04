<script lang="ts">
  import '../app.pcss';
  import Navbar from './components/navbar.svelte';
</script>

<main class="min-h-screen p-4">
  <slot></slot>
</main>
<Navbar />

<style lang="postcss">
  :global(body) {
    @apply mx-auto min-h-screen max-w-screen-sm;
  }
</style>
