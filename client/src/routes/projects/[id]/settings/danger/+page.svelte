<script lang="ts">
	import { ApiService } from '$lib/api';
	import Header from '../Header.svelte';
	import HeaderDescription from '../HeaderDescription.svelte';
	import SubHeader from '../SubHeader.svelte';
	import { project } from '../../store';
	let secretKey = '';

	console.log($project);
</script>

<div class="">
	<Header>Danger</Header>
	<HeaderDescription>Manage very sensitive settings about your project</HeaderDescription>
	<SubHeader>Keys</SubHeader>
	<HeaderDescription>See and manage keys</HeaderDescription>
	<div class="mb-6 rounded-md border p-5">
		<h3 class="font-semibold">Public key</h3>
		<p class="text-text-dimmed text-sm">This key is like the username to connect to this project</p>
		<button
			class="mt-4 cursor-pointer"
			on:click={() => {
				navigator.clipboard.writeText($project.id);
			}}
		>
			{$project.id}
		</button>
	</div>
	<div class="mb-6 rounded-md border bg-red-100 p-5 text-red-900">
		<h3 class="font-semibold">Secret key</h3>
		<p class="text-sm text-red-700">This key will only be shown once when it is created</p>
		{#if secretKey}
			<button
				class="mt-4 cursor-pointer"
				on:click={() => {
					navigator.clipboard.writeText($project.id);
				}}
			>
				{secretKey}
			</button>
		{:else}
			<button
				class="btn mt-4 cursor-pointer bg-red-800 text-red-50 hover:bg-red-900"
				on:click={async () => {
					const res = await ApiService.generateNewSecretKey($project.id);
					console.log(res);
					secretKey = res;
				}}
			>
				Generate new secret key
			</button>
		{/if}
	</div>
</div>
