<script lang="ts">
	import { ApiService } from '$lib/api';
	let name = '';
	let error = '';

	async function handleCreate() {
		try {
			const res = await ApiService.createNewProject(name);
			console.log(res);
			if (res.id) {
				window.location.href = `/projects/${res.id}`;
			}
		} catch (error: any) {
			error = error.message;
		}
	}
</script>

<h2 class="mb-3 text-4xl font-bold">Lets create a new project</h2>

<div class="">
	<label for="new-project-name" class="text-text-dimmed block">
		What should we call your project?
	</label>
	<input
		bind:value={name}
		type="text"
		id="new-project-name"
		name="new-project-name"
		placeholder="my project name..."
		class="focus:border-primary rounded-md border border-black px-2 py-1 text-sm outline-none"
	/>
</div>

{#if error}
	<div class="rounded-md bg-red-100 p-5 text-red-900">{error}</div>
{/if}

<button disabled={!name} on:click={handleCreate} class="btn my-4">Create new project</button>
