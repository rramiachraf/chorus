<script lang="ts">
	import { createQuery } from '@tanstack/svelte-query'
	import { playSong, currentSong } from '../components/player/audio'
	import Section from '../components/Section.svelte'
	import { tracksList } from '../store'

	interface Album {
		name: string
		picture: number
		artist: string
		year: number
		songs: {
			id: number
			title: string
			track: number
			disc: number
		}[]
	}

	export let albumID: number

	const query = createQuery<Album>({
		queryKey: ['album', albumID],
		queryFn: () => fetch(`/api/album/${albumID}`).then(res => res.json())
	})

	function startPlaying(id: number) {
		playSong(id)
		tracksList.set([])
		$query.data.songs.forEach(({ id }) => {
			tracksList.update(ids => [...ids, id])
		})
		localStorage.setItem('tracksList', JSON.stringify($tracksList))
	}
</script>

<Section>
	<div id="container">
		{#if $query.isSuccess}
			<div id="info">
				<figure style="background-image: url(/api/picture/{$query.data.picture});" />
				<section>
					<h2>{$query.data.artist}</h2>
					<div>
						<h1>{$query.data.name}</h1>
						<small>{$query.data.year}</small>
					</div>
				</section>
			</div>

			<div id="songs">
				{#each $query.data.songs as { id, title, track }}
					<button
						class:playing={id === Number($currentSong)}
						class="song"
						on:click={() => startPlaying(id)}>
						{#if track}<small>{track}</small>{/if}
						<div>{title}</div>
					</button>
				{/each}
			</div>
		{/if}
	</div>
</Section>

<style>
	#container {
		display: grid;
		grid-template-columns: 300px 1fr;
		align-items: flex-start;
		gap: 30px;
	}

	figure {
		width: 300px;
		height: 300px;
		background-size: cover;
		border-radius: 10px;
		margin: 0;
		border: 1px solid var(--fifth-color);
	}

	section {
		margin-top: 10px;
		display: flex;
		flex-direction: column;
		gap: 8px;
		text-align: center;
	}

	section div {
		display: flex;
		flex-direction: column;
		gap: 2px;
	}

	h1,
	h2 {
		margin: 0;
	}

	h1 {
		color: #eee;
		font-size: 22px;
		text-transform: uppercase;
	}

	h2 {
		color: #ccc;
		font-size: 14px;
		font-weight: normal;
	}

	small {
		font-size: 14px;
		color: #ccc;
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

	.song {
		padding: 10px 20px;
		background-color: #111;
		cursor: pointer;
		color: #ddd;
		font-size: 14px;
		display: flex;
		gap: 10px;
		align-items: center;
		border-radius: 5px;
		border: 1px solid transparent;
	}

	.song div {
		font-weight: 500;
	}

	.song:focus {
		border: 1px solid var(--forth-color);
	}

	.playing {
		background-color: var(--third-color);
		border: 1px solid var(--fifth-color);
	}
</style>
