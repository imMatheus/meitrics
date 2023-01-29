<script lang="ts">
	import classNames from 'classnames';
	export let data: { id: string; pathname: string };
	import { project } from './store';

	$: $project = data;

	const links = [
		{
			link: '/',
			text: 'Overview'
		},
		{
			link: '/settings',
			text: 'Settings'
		}
	];
</script>

<div class="mb-10">
	<div class="flex">
		{#each links as link}
			<a
				href={`/projects/${data.id}${link.link}`}
				class={classNames(
					'border-b px-4 pt-2 pb-4 text-sm transition-colors',
					typeof window !== 'undefined' &&
						(window.location.pathname === `/projects/${data.id}${link.link}` ||
						(window.location.pathname.startsWith(`/projects/${data.id}${link.link}`) &&
							link.link !== `/`)
							? 'border-b-primary text-primary'
							: 'hover:border-b-text-lighten text-text-lighten')
				)}
			>
				{link.text}
			</a>
		{/each}
	</div>
	<div class="absolute left-0 right-0 -z-10 -translate-y-[1px] border-b" />
</div>

<slot />
