<script lang="ts">
	import { currentSong, playSong } from './player/audio'
	import { tracksList } from '../routes/store'

	export let id: number
	export let artist: string
	export let title: string

	function play(id: number) {
		tracksList.set([])
		playSong(id)
	}
</script>

<button on:click={() => play(id)}>
	<article class:playing={Number($currentSong) === id}>
		<small>{artist || 'Unknown Artist'}</small>
		<p>{title}</p>
	</article>
</button>

<style>
	button {
		border: none;
		background: none;
		text-align: left;
	}

	article {
		background-color: var(--third-color);
		padding: 10px 15px;
		border-radius: 5px;
		display: flex;
		flex-direction: column;
		gap: 3px;
		cursor: pointer;
		transition: 0.2s ease-in background-color;
	}

	.playing {
		background-color: var(--forth-color);
	}

	p {
		color: #eee;
		font-size: 14px;
		font-weight: 500;
	}

	small {
		font-size: 12px;
		color: #aaa;
	}
</style>
