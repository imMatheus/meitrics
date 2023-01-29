import type { Load } from '@sveltejs/kit';

export const load: Load = async ({ params, url, route }) => {
	console.log('hi hi hi hi hi hi hi hi hi hi hi hi');

	return { id: params.id, pathname: url.pathname };
};
