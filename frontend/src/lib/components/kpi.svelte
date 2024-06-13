<script lang="ts">
  export let title: string;
</script>

<div
  class="flex flex-col justify-between gap-2 rounded bg-gray-800 p-4 shadow-sm"
>
  <div class="text-base text-gray-300">{title}</div>
  <div class="text-primary text-2xl font-bold">
    <slot />
  </div>
</div>
