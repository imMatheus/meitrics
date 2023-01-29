<script lang="ts">
	import Header from '../Header.svelte';
	import SubHeader from '../SubHeader.svelte';
	import { project } from '../../store';

	console.log($project);
</script>

<div class="">
	<Header>Danger</Header>
	<SubHeader>Manage very sensitive settings about your project</SubHeader>
</div>
