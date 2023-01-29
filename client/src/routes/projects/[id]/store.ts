import { writable } from 'svelte/store';

export const project = writable<{ id: string; pathname: string }>({ id: '', pathname: '' });
