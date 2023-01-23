/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				primary: '#1E9DE7',
				'primary-dimmed': '#0B6EC5',
				bg: '#fff',
				'bg-lighten': '#FAFAFA',
				'bg-dimmed': '#EBEBEB',
				text: '#000',
				'text-lighten': '#1A1A1A',
				'text-dimmed': '#616161'
			}
		}
	},
	plugins: []
};
