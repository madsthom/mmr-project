<script lang="ts">
  import { Button } from '$lib/components/ui/button';
  import { page } from '$app/stores';
  import { cn } from '$lib/utils';

  export let path: string;
  export let isPrimary = false;
</script>

<Button
  variant={isPrimary ? 'default' : 'ghost'}
  href={path}
  class={cn('h-16 min-w-16 flex-1', {
    'text-primary hover:text-primary':
      !isPrimary && path === $page.url.pathname,
  })}
>
  <slot />
</Button>
