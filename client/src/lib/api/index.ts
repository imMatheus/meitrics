import type { Project, Log } from '$lib/types';
import axios from 'axios';

axios.defaults.baseURL = 'http://localhost:3000';

export const ApiService = {
	getProjects: async () => {
		console.log('hej');
		const { data } = await axios.get<Project[]>('/projects');

		return data;
	},
	getProjectById: async (id: string) => {
		console.log('get Project by id');
		const { data } = await axios.get<Project>(`/projects/${id}`);

		return data;
	},
	GetLogsForProject: async (id: string) => {
		console.log('get Project by id');
		const { data } = await axios.get<Log[]>(`/projects/${id}/logs`);

		return data;
	}
};
