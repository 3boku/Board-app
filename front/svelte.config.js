import adapterAuto from '@sveltejs/adapter-auto';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// ...
	kit: {
		adapter: adapterAuto({
			// ...
		})
	}
};

export default config;
