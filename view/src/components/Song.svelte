<script lang="ts">
	import { currentSong, playSong } from './player/audio'
	import { tracksList } from '../store'

	export let id: number
	export let artist: string
	export let title: string
	export let picture: number

	function play(id: number) {
		tracksList.set([])
		playSong(id)
	}
</script>

<button on:click={() => play(id)}>
	<article class:playing={Number($currentSong) === id}>
		<figure style:background-image="url('/api/picture/{picture}')" />
		<div>
			<small>{artist || 'Unknown Artist'}</small>
			<p>{title}</p>
		</div>
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
		padding: 8px 10px;
		border-radius: 5px;
		gap: 15px;
		display: flex;
		cursor: pointer;
		align-items: center;
		transition: 0.2s ease-in background-color;
		border: 1px solid transparent;
	}

	figure {
		background-color: var(--second-color);
		background-size: cover;
		border: 1px solid var(--fifth-color);
		width: 40px;
		height: 40px;
		padding: 0;
		margin: 0;
		border-radius: 5px;
	}

	div {
		display: flex;
		flex-direction: column;
		gap: 3px;
	}

	.playing {
		background-color: var(--forth-color);
		border: 1px solid var(--fifth-color);
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
