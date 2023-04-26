<script lang="ts">
	import { onMount } from 'svelte'
	import Info from './Info.svelte'
	import Controls from './Controls.svelte'
	import Options from './Options.svelte'
	import { currentSong, loadLastSong, songError } from './audio'

	onMount(() => {
		loadLastSong()
	})
</script>

<main class:noSong={$currentSong === undefined}>
	{#if $songError}
		<button class="alert" on:click={() => songError.set(false)}>
			Can't play the current song.
		</button>
	{/if}
	<Info />
	<Controls />
	<Options />
</main>

<style>
	main {
		background-color: var(--third-color);
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		border-top: 1px solid var(--fifth-color);
		padding: 0 20px;
	}

	.noSong {
		display: none;
	}

	.alert {
		position: absolute;
		color: white;
		right: 0;
		top: 0;
		background-color: var(--main-color);
		padding: 10px;
		box-sizing: border-box;
		width: 100vw;
		font-size: 14px;
		text-align: center;
		z-index: 300;
		border: none;
		cursor: pointer;
		border-bottom-left-radius: 5px;
		border-bottom-right-radius: 5px;
	}
</style>
