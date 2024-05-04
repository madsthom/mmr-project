<script lang="ts">
  import Button from '../../lib/components/ui/button/button.svelte';
  import { Plus } from 'lucide-svelte';
</script>

<nav
  class="bg-card absolute bottom-0 left-0 right-0 flex h-16 flex-row items-center justify-center border-t"
>
  <Button variant="link" class="flex-[1.5]" href="/submit"><Plus /></Button>
</nav>
