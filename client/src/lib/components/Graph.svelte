<script lang="ts">
	import { scaleLinear } from 'd3-scale';

	let points: { x: number; y: number }[] = [];

	for (let i = 0; i < 300; i += 1) {
		for (let j = 0; j <= 1; j++) {
			points.push({ x: i, y: Math.sin(Math.random() * 2 * Math.PI) + 10 });
			// f.push({ x: i, y: (f[(i + 1) * j]?.y || 0) + Math.random() * 3 || j });
		}
	}

	const yTicks = [0, 4, 8, 12, 16, 20];

	const delta = Math.floor(points.length / 12);
	let xTicks: number[] = [];
	for (let i = 0; i < points.length; i = i + delta) {
		xTicks.push(points[i]?.x || 0);
	}
	const padding = { top: 20, right: 15, bottom: 20, left: 25 };

	let width = 500;
	let height = 200;

	$: minX = points[0]?.x || 0;
	$: maxX = points[points.length - 1]?.x || 0;
	$: xScale = scaleLinear()
		.domain([minX, maxX])
		.range([padding.left, width - padding.right]);

	$: yScale = scaleLinear()
		.domain([Math.min.apply(null, yTicks), Math.max.apply(null, yTicks)])
		.range([height - padding.bottom, padding.top]);

	$: path = `M${points.map((p) => `${xScale(p.x)},${yScale(p.y)}`).join('L')}`;
	$: area = `${path}L${xScale(maxX)},${yScale(0)}L${xScale(minX)},${yScale(0)}Z`;

	function formatMobile(tick: number) {
		return "'" + tick.toString().slice(-2);
	}
</script>

<div class="chart" bind:clientWidth={width} bind:clientHeight={height}>
	<svg>
		<!-- y axis -->
		<g class="axis y-axis" transform="translate(0, {padding.top})">
			{#each yTicks as tick}
				<g class="tick tick-{tick}" transform="translate(0, {yScale(tick) - padding.bottom})">
					<line x2="100%" />
					<text y="-4">{tick}</text>
				</g>
			{/each}
		</g>

		<!-- x axis -->
		<g class="axis x-axis">
			{#each xTicks as tick}
				<g class="tick tick-{tick}" transform="translate({xScale(tick)},{height})">
					<line y1="-{height}" y2="-{padding.bottom}" x1="0" x2="0" />
					<text y="-2">{width > 380 ? tick : formatMobile(tick)}</text>
				</g>
			{/each}
		</g>

		<!-- data -->
		<path class="path-area" d={area} />
		<path class="path-line" d={path} />
	</svg>
</div>

<style>
	.chart,
	h2,
	p {
		width: 100%;
	}

	svg {
		position: relative;
		width: 100%;
		height: 200px;
		overflow: visible;
	}

	.tick {
		font-size: 0.725em;
		font-weight: 200;
	}

	.tick line {
		stroke: #aaa;
		stroke-dasharray: 1;
	}

	.tick text {
		fill: #666;
		text-anchor: start;
	}

	.tick.tick-0 line {
		stroke-dasharray: 1;
	}

	.x-axis .tick text {
		text-anchor: middle;
	}

	.path-line {
		fill: none;

		stroke-linejoin: round;
		stroke-linecap: round;
		stroke-width: 1;

		@apply stroke-primary;
	}

	.path-area {
		@apply fill-primary/20;
	}
</style>
