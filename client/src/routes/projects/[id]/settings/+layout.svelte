<script lang="ts">
	import classNames from 'classnames';
	import { project } from '../store';
	// export let data: { id: string; pathname: string };

	// $: $project = data;

	const links = [
		{
			link: '/settings',
			text: 'General'
		},
		{
			link: '/settings/team',
			text: 'Team'
		},
		{
			link: '/settings/danger',
			text: 'Danger'
		}
	];
	console.log(typeof window !== 'undefined' && window.location);
</script>

<div class="flex flex-col gap-4 lg:flex-row">
	<div class="-mt-1 hidden w-1/5 flex-col text-sm lg:flex">
		{#each links as link}
			<a
				href={`/projects/${$project.id}${link.link}`}
				class={classNames(
					'py-2',
					$project.pathname === `/projects/${$project.id}${link.link}`
						? 'text-text-lighten font-semibold'
						: 'text-text-dimmed hover:text-text-lighten'
				)}
				>{link.text}
			</a>
		{/each}
	</div>
	<div class="flex-1">
		<slot />
	</div>
</div>
