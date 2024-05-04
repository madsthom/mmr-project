<script lang="ts">
  import * as Form from '$lib/components/ui/form';
  import {
    superForm,
    type Infer,
    type SuperValidated,
  } from 'sveltekit-superforms';
  import { zodClient } from 'sveltekit-superforms/adapters';
  import { matchSchema } from '../match-schema';
  import type { MatchForm, MatchSchema } from '../match-schema';
  import { Input } from '$lib/components/ui/input';
  import MatchFormField from './match-form-field.svelte';
  import { Button } from '$lib/components/ui/button';
  import MatchCard from '$lib/components/match-card.svelte';
  import { fade, slide, fly } from 'svelte/transition';

  export let data: SuperValidated<Infer<MatchSchema>>;

  let isValid = false;

  const form = superForm(data, {
    validators: zodClient(matchSchema),
    dataType: 'json',
  });

  const { form: formData, enhance } = form;

  let loosingTeam: 'team1' | 'team2' | null = null;
  $: loosingTeam =
    $formData.team1.score === 10
      ? 'team2'
      : $formData.team2.score === 10
        ? 'team1'
        : null;

  const setTeam1AsWinner = () => {
    $formData.team1.score = 10;
    $formData.team2.score = -1;
  };
  const setTeam2AsWinner = () => {
    $formData.team2.score = 10;
    $formData.team1.score = -1;
  };

  $: isMatchCardVisible =
    loosingTeam !== null &&
    $formData.team1.score !== -1 &&
    $formData.team2.score !== -1;

  $: allInitialsFilled =
    $formData.team1.member1 != '' &&
    $formData.team1.member2 != '' &&
    $formData.team2.member1 != '' &&
    $formData.team2.member2 != '';
</script>

<form method="POST" use:enhance>
  <div class="flex flex-col gap-2">
    <div class="flex flex-row gap-4">
      <div class="flex flex-1 flex-col">
        <h3 class="mb-2 text-center text-2xl">Team 1</h3>
        <MatchFormField
          {form}
          label="You"
          name="team1.member1"
          bind:value={$formData.team1.member1}
          placeholder="Enter initials"
        />
        <MatchFormField
          {form}
          label="Your teammate"
          name="team1.member2"
          bind:value={$formData.team1.member2}
          placeholder="Enter initials"
        />
      </div>
      <div class="flex-s bg-border min-h-full w-px"></div>
      <div class="flex-1">
        <h3 class="mb-2 text-center text-2xl">Team 2</h3>
        <MatchFormField
          {form}
          label="Opponent 1"
          name="team2.member1"
          bind:value={$formData.team2.member1}
          placeholder="Enter initials"
        />
        <MatchFormField
          {form}
          label="Opponent 2"
          name="team2.member2"
          bind:value={$formData.team2.member2}
          placeholder="Enter initials"
        />
      </div>
    </div>
    {#if allInitialsFilled}
      <div class="flex flex-col gap-4" transition:fade>
        <h2 class="text-center text-4xl">Who won?</h2>
        <div class="flex flex-row gap-4">
          <Button
            on:click={setTeam1AsWinner}
            class="flex-1"
            variant="default"
            disabled={$formData.team1.score === 10}>We won &nbsp; 🎉</Button
          >
          <div class="flex-s bg-border min-h-full w-px"></div>
          <Button
            on:click={setTeam2AsWinner}
            class="flex-1"
            variant="destructive"
            disabled={$formData.team2.score === 10}>They won &nbsp; 😓</Button
          >
        </div>
      </div>
    {/if}
    {#if loosingTeam}
      <div class="flex flex-col gap-4" transition:fade>
        <h2 class="text-center text-4xl">
          What was {loosingTeam === 'team1' ? 'your' : 'their'} score?
        </h2>
        <div class="grid grid-cols-5 gap-2">
          {#each Array.from({ length: 10 }, (_, i) => i) as score}
            <Button
              variant={$formData[loosingTeam].score === score
                ? 'default'
                : 'outline'}
              on:click={() => {
                $formData[loosingTeam].score = score;
              }}
            >
              {score === 0 ? '🥚' : score}
            </Button>
          {/each}
        </div>
      </div>
    {/if}
    {#if isMatchCardVisible}
      <div class="flex flex-col gap-4" transition:fade>
        <h2 class="text-center text-4xl">Submit?</h2>
        <MatchCard match={$formData} />
        <Button>Submit match</Button>
      </div>
    {/if}
  </div>
</form>
