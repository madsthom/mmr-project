<script>
  import PageTitle from '$lib/components/page-title.svelte';
  import Button from '$lib/components/ui/button/button.svelte';
  import * as Card from '$lib/components/ui/card';
  import { Info } from 'lucide-svelte';
</script>

<div class="flex flex-col gap-3">
  <PageTitle>Profile</PageTitle>

  <Card.Root class="flex flex-col gap-3 border-white p-4">
    <div class="flex space-x-2 text-white">
      <Info /><span>You have not claimed a player yet</span>
    </div>
    <Button href="/profile/claim">Claim a player</Button>
  </Card.Root>

  <form method="post" action="/signout" class="contents">
    <Button type="submit" variant="secondary">Logout</Button>
  </form>
</div>
