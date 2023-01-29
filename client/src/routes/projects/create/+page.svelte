<script lang="ts">
	let name = '';
	import { ApiService } from '$lib/api';

	async function handleCreate() {
		const res = await ApiService.createNewProject(name);
		console.log(res);
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

<button disabled={!name} on:click={handleCreate} class="btn my-4">Create new project</button>
