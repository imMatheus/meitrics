<script lang="ts">
	import { ApiService } from '$lib/api';
	import LoadSpinner from '$lib/components/LoadSpinner.svelte';

	const pending = ApiService.getProjects();
	console.log(pending);
</script>

{#await pending}
	<LoadSpinner />
{:then projects}
	<div class="flex justify-between gap-3 items-center">
		<h3 class="font-bold text-3xl mb-4">Your projects</h3>
		<a href="/projects/create" class="btn">Create a new</a>
	</div>
	<div class="grid gap-5 sm:grid-cols-2 lg:grid-cols-3 ">
		{#each projects as project}
			<a
				href={`/projects/${project.id}`}
				class="rounded-md group cursor-pointer overflow-hidden border border-text"
			>
				<div
					class="h-20 group-hover:bg-primary border-b transition-colors border-text bg-rose-500"
				/>
				<div class="p-5">
					<h4 class="font-bold mb-2 text-lg">{project.name}</h4>
					<p class="text-text-dimmed text-sm font-medium">23k/s reads</p>
					<p class="text-text-dimmed text-xs">Created 12 days ago</p>
				</div>
			</a>
		{/each}
	</div>
{:catch error}
	<p>shit...</p>
{/await}
