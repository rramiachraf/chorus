<script>
	import { createEventDispatcher } from 'svelte'
	import { isSearching } from '../store'

	export let value = ''

	const dispatch = createEventDispatcher()
</script>

<input
	bind:value
	on:keyup={() => dispatch('input')}
	on:focus={() => isSearching.set(true)}
	on:blur={()=> isSearching.set(false)}
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
