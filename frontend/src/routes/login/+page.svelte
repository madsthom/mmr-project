<script lang="ts">
  import { ThemeSupa } from '@supabase/auth-ui-shared';
  import { Auth } from '@supabase/auth-ui-svelte';

  export let data;
</script>

<div class="mx-auto mb-16 max-w-screen-sm overflow-auto p-4">
  <div class="row flex-center flex">
    <div class="col-6 form-widget">
      <Auth
        supabaseClient={data.supabase}
        view="sign_in"
        redirectTo={`${data.url}/auth/callback`}
        appearance={{ theme: ThemeSupa, style: { input: 'color: #fff' } }}
      />
    </div>
  </div>
</div>
