<script lang="ts">
  import {
    HeatmapChart,
    ScaleTypes,
    type AreaChartOptions,
    type HeatmapChartOptions,
  } from '@carbon/charts-svelte';
  import '@carbon/charts-svelte/styles.css';

  export let data: Array<{
    dayOfWeek: number;
    hourOfDay: number;
    count: number;
  }>;

  const hourOfDayFormatter = new Intl.DateTimeFormat(undefined, {
    hour: 'numeric',
    hour12: false,
  });
  const dayOfWeekFormatter = new Intl.DateTimeFormat(undefined, {
    weekday: 'long',
  });

  const dayOfWeekLabels = Array.from(
    { length: 7 },
    (_, i) => {
      const date = new Date();
      date.setDate(date.getDate() - date.getDay() + i + 1); // we + 1 to make 0 = monday
      return dayOfWeekFormatter.format(date);
    }
    // dayOfWeekFormatter.format(new Date(2024, 8, i)) // Month that starts on a Monday (could be any - just for sorting)
  );
  const hourOfDayLabels = Array.from({ length: 24 }, (_, i) =>
    hourOfDayFormatter.format(new Date(0, 0, 0, i))
  );

  console.table(data);
  console.log(dayOfWeekLabels);

  const mappedChartData =
    data?.map((stat) => ({
      day: dayOfWeekLabels[stat.dayOfWeek + 1], // we + 1 to make 0 = monday
      hour: hourOfDayLabels[stat.hourOfDay],
      value: stat.count,
    })) ?? [];

  // https://github.com/carbon-design-system/carbon-charts/pull/1846
  const chartOptions: HeatmapChartOptions & { axes: AreaChartOptions['axes'] } =
    {
      theme: 'g100',
      height: '320px',
      axes: {
        bottom: {
          title: 'Hour of day',
          mapsTo: 'hour',
          scaleType: ScaleTypes.LABELS,
          domain: hourOfDayLabels,
        },
        left: {
          title: 'Day of week',
          mapsTo: 'day',
          scaleType: ScaleTypes.LABELS,
          domain: dayOfWeekLabels.toReversed(), // Needed to make monday appear on top
        },
      },
      heatmap: {
        colorLegend: {
          title: 'Frequency',
          type: 'linear',
        },
      },
    };
</script>

<!-- Added margin on bottom to make sure heatmap legend is visible -->
<div class="mb-2">
  <HeatmapChart data={mappedChartData} options={chartOptions} />
</div>
