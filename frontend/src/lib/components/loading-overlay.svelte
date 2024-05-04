<script lang="ts">
  import * as Dialog from '$lib/components/ui/dialog';
  import { LoaderCircle } from 'lucide-svelte';
  export let isLoading: boolean;
</script>

<Dialog.Root
  bind:open={isLoading}
  closeOnEscape={false}
  closeOnOutsideClick={false}
>
  <Dialog.Content
    class="loading-dialog flex flex-col items-center justify-center"
  >
    Uploading match result
    <LoaderCircle class="animate-spin" />
  </Dialog.Content>
</Dialog.Root>

<style>
  :global(.loading-dialog > button[data-dialog-close]) {
    display: none;
  }
</style>
