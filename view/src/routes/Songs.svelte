<script lang="ts">
	import { onMount } from 'svelte'
	import { currentSong, playSong } from '../components/player/audio'
	import Search from '../components/Search.svelte'
	import { songs } from './store'

	let visibleSongs = []
	let search = ''

	onMount(async () => {
		if ($songs.length === 0) {
			const res = await fetch('/api/songs')
			songs.set(await res.json())
		}
		visibleSongs = $songs
	})

	function searchSong() {
		const s = search.toLowerCase()
		if (s === '') {
			visibleSongs = $songs
			return
		}

		visibleSongs = $songs.filter(song => {
			return (
				song.title.toLowerCase().includes(s) ||
				song.artist.toLowerCase().includes(s) ||
				song.album && song.album.toLowerCase().includes(s)
			)
		})
	}
</script>

<section>
	<Search on:input={searchSong} bind:value={search} />
	<div>
		{#each visibleSongs as { id, title, artist }}
			<button on:click={() => playSong(id)}>
				<article {id} class:playing={Number($currentSong) === id}>
					<small>{artist || 'Unknown Artist'}</small>
					<p>{title}</p>
				</article>
			</button>
		{/each}
	</div>
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		padding: 20px;
		gap: 20px;
		overflow-y: auto;
	}

	div {
		display: flex;
		flex-direction: column;
		gap: 5px;
	}

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
