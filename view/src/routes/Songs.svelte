<script lang="ts">
	import { createQuery } from '@tanstack/svelte-query'
	import Search from '../components/Search.svelte'
	import Section from '../components/Section.svelte'
	import SongSnippet from '../components/Song.svelte'

	interface Song {
		id: number
		title: string
		artist: string
		album: string
	}

	const query = createQuery<Song[]>({
		queryKey: ['songs'],
		queryFn: () => fetch('/api/songs').then(res => res.json()),
	})

	$: visibleSongs = $query.data
	let search = ''

	function searchSong() {
		const s = search.toLowerCase()
		if (s === '') {
			visibleSongs = $query.data
			return
		}

		visibleSongs = $query.data.filter(song => {
			return (
				song.title.toLowerCase().includes(s) ||
				song.artist.toLowerCase().includes(s) ||
				song.album && song.album.toLowerCase().includes(s)
			)
		})
	}
</script>

<Section>
	{#if $query.isSuccess}
		<div id="container">
			<Search on:input={searchSong} bind:value={search} />
			<div id="songs">
				{#each visibleSongs as { id, title, artist }}
					<SongSnippet {id} {title} {artist} />
				{/each}
			</div>
		</div>
	{/if}
</Section>

<style>
	#container {
		display: flex;
		flex-direction: column;
		gap: 20px;
	}

	#songs {
		display: flex;
		flex-direction: column;
		gap: 5px;
	}
</style>
