<script lang="ts">
	import { onMount } from 'svelte'
	import Info from './Info.svelte'
	import Controls from './Controls.svelte'
	import Options from './Options.svelte'
	import { currentSong, loadLastSong } from './audio'

	onMount(() => {
		loadLastSong()
	})
</script>

<main class:noSong={$currentSong === undefined}>
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
</style>
