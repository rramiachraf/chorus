<script lang="ts">
	import { createEventDispatcher } from 'svelte'

	export let value: string
	export let disabled = false

	const dispatch = createEventDispatcher()
</script>

<button {disabled} on:click={() => dispatch('click')}>{value}</button>

<style>
	button {
		background-color: var(--main-color);
		border-radius: 5px;
		color: white;
		cursor: pointer;
		border: 2px solid transparent;
		padding: 6px 20px;
		font-size: 14px;
		font-weight: 500;
		box-shadow: 0 1px 3px var(--fifth-color);
		transition: .3s ease-out background-color;
		align-self: var(--align-self);
	}

	button:hover:enabled {
		background-color: transparent;
		border: 2px solid var(--main-color);
		color: var(--main-color);
	}

	button:disabled {
		opacity: 0.7;
	}
</style>
