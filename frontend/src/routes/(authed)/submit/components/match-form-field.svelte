<script lang="ts">
  import * as Form from '$lib/components/ui/form';
  import { Input } from '$lib/components/ui/input';
  import type { FormPath, Infer, SuperForm } from 'sveltekit-superforms';
  import type { MatchSchema } from '../match-schema';

  export let form: SuperForm<Infer<MatchSchema>>;
  export let name: FormPath<Infer<MatchSchema>>;
  export let value: string;
  export let label: string;
  export let placeholder = '';
</script>

<Form.Field {form} {name}>
  <Form.Control let:attrs>
    <Form.Label>{label}</Form.Label>
    <Input
      {...attrs}
      bind:value
      class="lowercase placeholder:normal-case"
      {placeholder}
    />
  </Form.Control>
  <Form.FieldErrors />
</Form.Field>
