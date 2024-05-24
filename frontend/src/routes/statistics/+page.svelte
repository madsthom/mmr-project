<script lang="ts">
  import PageTitle from '$lib/components/page-title.svelte';
  import {
    LineChart,
    ScaleTypes,
    type ChartTabularData,
    type LineChartOptions,
  } from '@carbon/charts-svelte';
  import '@carbon/charts-svelte/styles.css';

  const data: ChartTabularData = [
    {
      player: 'thj',
      date: '2021-01-01',
      rating: 1000,
    },
    {
      player: 'thj',
      date: '2021-01-02',
      rating: 1010,
    },
    { player: 'bhs', date: '2021-01-02', rating: 1100 },
    { player: 'bhs', date: '2021-01-05', rating: 1090 },
    { player: 'bhs', date: '2021-01-06', rating: 1095 },
  ];
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
<LineChart {data} {options} />
