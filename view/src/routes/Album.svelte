<script lang="ts">
	import { onMount } from "svelte"
	import { playSong, currentSong } from "../components/player/audio"

	export let albumID: number
	let name: string
	let picture: string
	let artist: string
	let songs = []

	onMount(async () => {
		const res = await fetch(`/api/album/${albumID}`)
		const json = await res.json()
		name = json.name
		picture = `/api/picture/${json.picture}`
		songs = json.songs
		artist = json.artist
	})
</script>

<section>
	<div id="info">
		<figure style="background-image: url({picture});" />
		<h1>{name}</h1>
		<h2>{artist}</h2>
	</div>
	<div id="songs">
		{#each songs as { id, title, track }}
			<div class:playing={id === $currentSong} id="song" on:click={() => playSong(id)}>
				{#if track}<small>{track}</small>{/if}
				<div>{title}</div>
			</div>
		{/each}
	</div>
</section>

<style>
	section {
		overflow-y: auto;
		padding: 20px;
		display: grid;
		grid-template-columns: 300px 1fr;
		align-items: flex-start;
		gap: 30px;
	}

	figure {
		width: 300px;
		height: 300px;
		background-size: cover;
		border-radius: 5px;
		margin: 0;
	}

	h1,
	h2 {
		text-align: center;
		margin: 0;
		padding: 5px;
	}

	h1 {
		color: #eee;
		font-size: 25px;
	}

	h2 {
		color: #ccc;
		font-size: 14px;
		font-weight: normal;
	}

	#info {
		position: sticky;
		top: 0;
		left: 0;
		display: flex;
		flex-direction: column;
		gap: 5px;
	}

	#songs {
		flex-grow: 1;
		display: flex;
		flex-direction: column;
		gap: 5px;
		padding-bottom: 20px;
	}

	#song {
		padding: 10px 20px;
		background-color: #111;
		cursor: pointer;
		color: #ddd;
		font-size: 14px;
		display: flex;
		gap: 10px;
		align-items: center;
		border-radius: 5px;
	}
	
	#song div {
		font-weight: 500;
	}

	.playing {
		border: 1px solid #555;
	}
</style>
