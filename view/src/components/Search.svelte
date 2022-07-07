<script>
	import { createEventDispatcher } from 'svelte'

	export let value = ''

	const dispatch = createEventDispatcher()
</script>

<input
	bind:value
	on:input={() => dispatch('input')}
	type="text"
	placeholder="Search..."
/>

<style>
	input {
		padding: 10px;
		border-radius: 3px;
		border: 1px solid var(--fifth-color);
		text-align: center;
		background: none;
		color: white;
		transition: 0.3s ease border;
		font-size: 14px;
	}

	input:focus {
		border: 1px solid #bbb;
	}
</style>
