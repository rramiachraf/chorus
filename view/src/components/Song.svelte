<script lang="ts">
	import { currentSong, playSong } from './player/audio'
	import { tracksList } from '../store'

	export let id: number
	export let artist: string
	export let title: string
	export let picture: number
	export let mime: string

	function play(id: number) {
		tracksList.set([])
		playSong(id)
	}
</script>

<button on:click={() => play(id)}>
	<article class:playing={Number($currentSong) === id}>
		<div id="track">
			<figure style:background-image="url('/api/picture/{picture}')" />
			<div id="text">
				<small>{artist || 'Unknown Artist'}</small>
				<p>{title}</p>
			</div>
		</div>
		{#if mime}
			<div id="ext">{mime}</div>
		{/if}
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
		padding: 8px 15px;
		border-radius: 5px;
		display: flex;
		cursor: pointer;
		align-items: center;
		justify-content: space-between;
		transition: 0.2s ease-in background-color;
		border: 1px solid transparent;
	}

	#track {
		display: flex;
		gap: 15px;
		align-items: center;
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

	#text {
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

	#ext {
		background-color: var(--second-color);
		font-size: 11px;
		justify-self: end;
		border-radius: 100px;
		padding: 2px 8px;
		color: white;
		border: 1px solid var(--fifth-color);
	}
</style>
