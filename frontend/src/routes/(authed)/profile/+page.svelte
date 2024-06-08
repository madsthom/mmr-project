<script>
  import PageTitle from '$lib/components/page-title.svelte';
  import Button from '$lib/components/ui/button/button.svelte';
</script>

<PageTitle>Profile</PageTitle>

<form method="post" action="/signout" class="contents">
  <Button type="submit">Logout</Button>
</form>
