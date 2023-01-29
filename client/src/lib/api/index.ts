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
	getLogsForProject: async (id: string) => {
		console.log('get Project by id');
		const { data } = await axios.get<Log[]>(`/projects/${id}/logs`);

		return data;
	},
	createNewProject: async (name: string) => {
		console.log('create new project');
		const { data } = await axios.post<Project>(`/project`, { name });

		return data;
	},
	generateNewSecretKey: async (id: string) => {
		console.log('create new secret key');
		const { data } = await axios.patch<string>(`/projects/${id}/update-secret-key`);
		console.log(data);

		return data;
	}
};
