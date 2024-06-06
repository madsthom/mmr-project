<script lang="ts">
  import PageTitle from '$lib/components/page-title.svelte';
  import * as Form from '$lib/components/ui/form';
  import { Input } from '$lib/components/ui/input';
  import { superForm } from 'sveltekit-superforms';
  import { zodClient } from 'sveltekit-superforms/adapters';
  import type { PageServerData } from './$types';
  import { createUserSchema } from './schema';

  export let data: PageServerData;

  const form = superForm(data.form, {
    validators: zodClient(createUserSchema),
    dataType: 'json',
    delayMs: 500,
  });

  const { form: formData, enhance, submitting } = form;
</script>

<form method="post" use:enhance class="flex flex-col gap-1">
  <PageTitle>Create user</PageTitle>

  <Form.Field {form} name="name">
    <Form.Control let:attrs>
      <Form.Label>Initials</Form.Label>
      <Input
        {...attrs}
        bind:value={$formData.name}
        placeholder="Enter user's initials"
      />
    </Form.Control>
    <Form.FieldErrors />
  </Form.Field>
  <Form.Field {form} name="displayName">
    <Form.Control let:attrs>
      <Form.Label>Name</Form.Label>
      <Input
        {...attrs}
        bind:value={$formData.displayName}
        placeholder="Enter user's name"
      />
    </Form.Control>
    <Form.FieldErrors />
  </Form.Field>
  <Form.Button>Create user</Form.Button>
</form>
