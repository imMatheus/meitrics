<script lang="ts">
	import type { Log } from '$lib/types';
	export let logs: Log[];
	console.log(logs);

	import classNames from 'classnames';
</script>

<div class="relative min-w-0 max-md:!overflow-y-visible max-md:overflow-x-scroll">
	<table class="my-0 mx-5 min-w-full text-xs sm:mx-0 md:text-sm">
		<thead class="text-text-dimmed bg-bg-dimmed h-9 text-left">
			<tr>
				<th class="h-full font-normal">
					<div class="flex items-center px-2">Type</div>
				</th>
				<th class="h-full font-normal">
					<div class="flex items-center px-2">Url</div>
				</th>
				<th class="h-full font-normal">
					<div class="flex items-center px-2">IP address</div>
				</th>
				<th class="h-full font-normal">
					<div class="flex items-center px-2">Hostname</div>
				</th>
				<th class="h-full font-normal">
					<div class="flex items-center px-2">Message</div>
				</th>
				<th class="h-full font-normal">
					<div class="flex items-center px-2">Date</div>
				</th>
			</tr>
		</thead>
		<tbody class="divide-y border-x border-b">
			{#each logs as log}
				<tr class="h-11">
					<td class="relative h-full">
						<div
							class="peer flex h-full min-w-[0rem] max-w-lg items-center overflow-hidden text-ellipsis whitespace-nowrap px-2"
						>
							<div
								class={classNames(
									'rounded-full px-2 py-0.5 text-xs',
									log.type === 'error'
										? 'bg-red-100 text-red-900'
										: log.type === 'info'
										? 'bg-sky-100 text-sky-900'
										: log.type === 'warning'
										? 'bg-amber-100 text-amber-900'
										: 'bg-gray-100'
								)}
							>
								{log.type || 'other'}
							</div>
						</div>
					</td>

					<td class="relative h-full">
						<div
							class="peer flex h-full min-w-[0rem] max-w-lg items-center overflow-hidden text-ellipsis whitespace-nowrap px-2"
						>
							{log.url}
						</div>
					</td>
					<td class="relative h-full">
						<div
							class="peer flex h-full min-w-[0rem] max-w-lg items-center overflow-hidden text-ellipsis whitespace-nowrap px-2"
						>
							{log.ip}
						</div>
					</td>
					<td class="relative h-full">
						<div
							class="peer flex h-full min-w-[16rem] max-w-lg items-center overflow-hidden text-ellipsis whitespace-nowrap px-2"
						>
							{log.hostname}
						</div>
					</td>
					<td class="relative h-full">
						<div
							class="peer flex h-full min-w-[16rem] max-w-lg items-center overflow-hidden text-ellipsis whitespace-nowrap px-2"
						>
							{log.message}
						</div>
					</td>

					<td class="relative h-full">
						<div
							class="h-ful peer flex max-w-lg items-center overflow-hidden text-ellipsis whitespace-nowrap px-2"
						>
							{new Intl.DateTimeFormat(undefined, {
								dateStyle: 'medium'
								// timeStyle: 'medium'
							}).format(new Date(log.createdAt))}
						</div>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
