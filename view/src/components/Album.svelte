<script lang="ts">
	import { createEventDispatcher } from 'svelte'

	export let name: string
	export let picture: string
	export let artist!: string

	const dispatch = createEventDispatcher()
</script>

<div id="album" on:click={() => dispatch('click')}>
	<div id="picture" style="background-image: url('/api/picture/{picture}')" />
	{#if artist}<p id="artist">{artist}</p>{/if}
	<p id="name">{name}</p>
</div>

<style>
	#album {
		display: flex;
		flex-direction: column;
		gap: 5px;
		width: 100%;
		cursor: pointer;
		border-radius: 5px;
		transition: 0.3s ease background-color;
		padding-bottom: 5px;
	}

	#album:hover {
		background-color: #161616;
		opacity: 0.9;
	}

	p {
		margin: 0;
		padding: 0;
		text-align: center;
	}

	#picture {
		height: 200px;
		background-size: cover;
		border-radius: 5px;
	}

	#name {
		font-size: 15px;
		color: #ddd;
		font-weight: 600;
	}

	#artist {
		color: #bbb;
		font-size: 13px;
	}
</style>
