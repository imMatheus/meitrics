<script lang="ts">
	import { ApiService } from '$lib/api';

	import LogsTable from '$lib/components/LogsTable.svelte';
	import LoadSpinner from '$lib/components/LoadSpinner.svelte';
	import Graph from '$lib/components/Graph.svelte';
	import classNames from 'classnames';
	import { project } from './store';

	// https://dribbble.com/shots/18890801-Online-course-dashboard-Untitled-UI
	// https://dribbble.com/shots/16915378-UI-Details-Application-Logs-Server-Traffic-Data
	const pending = ApiService.getLogsForProject($project.id || '');
	console.log($project);
	const filters = ['All', 'Some', 'Maybe all'] as const;
	let selectedFilter: (typeof filters)[number] = 'All';
</script>

{#await pending}
	<LoadSpinner />
{:then value}
	<Graph />

	<h3 class="font-bold">Browse logs</h3>
	<p class="text-text-dimmed mb-2 text-sm">Filter all your logs just the way you want it</p>

	<div
		class="w-max mb-5 flex text-sm font-medium rounded-md overflow-hidden border border-text divide-x divide-text"
	>
		{#each filters as filter}
			<button
				class={classNames(
					'py-1.5 px-5 cursor-pointer transition-colors',
					filter === selectedFilter && 'bg-text text-bg'
				)}
				on:click={() => {
					selectedFilter = filter;
				}}
			>
				{filter}
			</button>
		{/each}
	</div>
	<LogsTable logs={value} />
{:catch error}
	<p>error</p>
{/await}
