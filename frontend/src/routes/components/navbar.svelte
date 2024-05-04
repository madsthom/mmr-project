<script lang="ts">
  import Button from '../../lib/components/ui/button/button.svelte';
  import { Plus, CircleUser, Home } from 'lucide-svelte';
  import NavbarNav from './navbar-nav.svelte';
</script>

<nav class="bg-card sticky bottom-0 left-0 right-0 border-t">
  <div
    class=" mx-auto flex max-w-screen-sm flex-row items-center justify-stretch"
  >
    <NavbarNav path="/">
      <Home />
    </NavbarNav>
    <NavbarNav isPrimary path="/submit"><Plus /></NavbarNav>
    <NavbarNav path="/profile"><CircleUser /></NavbarNav>
  </div>
</nav>
