<script lang="ts">
  import PageTitle from '$lib/components/page-title.svelte';
  import {
    LineChart,
    ScaleTypes,
    type ChartTabularData,
    type LineChartOptions,
  } from '@carbon/charts-svelte';
  import '@carbon/charts-svelte/styles.css';
  import type { PageData } from './$types';

  export let data: PageData;

  const chartData: ChartTabularData =
    data.statistics?.map((stat) => ({
      player: stat.name,
      date: stat.date,
      rating: stat.mmr,
    })) ?? [];

  const options: LineChartOptions = {
    theme: 'g100',
    curve: 'curveMonotoneX',
    data: {
      groupMapsTo: 'player',
      loading: false,
    },
    axes: {
      left: {
        title: 'Rating',
        mapsTo: 'rating',
        scaleType: ScaleTypes.LINEAR,
        includeZero: false,
      },
      bottom: {
        title: 'Time',
        mapsTo: 'date',
        scaleType: ScaleTypes.TIME,
      },
    },
  };
</script>

<PageTitle>Statistics</PageTitle>
<div class="container px-16">
  <LineChart data={chartData} {options} />
</div>
