<script lang="ts">
	import '../app.css';
	import Navbar from '$lib/components/Navbar.svelte';
</script>

<Navbar />
<div class="px-5 py-6 md:py-10">
	<div class="mx-auto max-w-7xl">
		<slot />
	</div>
</div>
