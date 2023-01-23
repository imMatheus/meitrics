<script lang="ts">
	import { ApiService } from '$lib/api';
	export let data: { id: string };

	// https://dribbble.com/shots/18890801-Online-course-dashboard-Untitled-UI
	const pending = ApiService.getProjectById(data.id || '');
	console.log(data);
</script>

{#await pending}
	<p>loading...</p>
{:then value}
	<p>data</p>
	<!-- pending was fulfilled -->
{:catch error}
	<p>error</p>
	<!-- pending was rejected -->
{/await}
