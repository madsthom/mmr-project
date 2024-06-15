<script lang="ts">
  import {
    HeatmapChart,
    ScaleTypes,
    type AreaChartOptions,
    type HeatmapChartOptions,
  } from '@carbon/charts-svelte';

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

  const dayOfWeekLabels = Array.from({ length: 7 }, (_, i) => {
    const date = new Date();
    date.setDate(date.getDate() - date.getDay() + i + 1); // we + 1 to make 0 = monday
    return dayOfWeekFormatter.format(date);
  });
  const hourOfDayLabels = Array.from({ length: 24 }, (_, i) =>
    hourOfDayFormatter.format(new Date(0, 0, 0, i))
  );

  const mappedChartData =
    data?.map((stat) => ({
      day: dayOfWeekLabels[(stat.dayOfWeek + 7 - 1) % 7], // - 1 to make 0 = monday, and +7 % 7 to handle when dayOfWeek == 0
      hour: hourOfDayLabels[stat.hourOfDay],
      value: stat.count * Math.random() * 100,
    })) ?? [];

  // https://github.com/carbon-design-system/carbon-charts/pull/1846
  const chartOptions: HeatmapChartOptions & { axes: AreaChartOptions['axes'] } =
    {
      theme: 'g100',
      height: '300px',
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
          type: 'quantize',
        },
      },
      color: {
        gradient: {
          colors: Array.from({ length: 10 }).map(
            (_, i) => `hsl(var(--primary) / ${(i + 1) / 10})`
          ),
        },
      },
    };
</script>

<!-- Added padding on bottom to make sure heatmap legend is visible -->
<div class="pb-2">
  <HeatmapChart data={mappedChartData} options={chartOptions} />
</div>
