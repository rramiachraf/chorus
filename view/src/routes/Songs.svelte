<script lang="ts">
	import { createInfiniteQuery } from '@tanstack/svelte-query'
	import Section from '../components/InfiniteSection.svelte'
	import SongSnippet from '../components/Song.svelte'
	import Button from '../components/Button.svelte'
	import { getNextPageParam } from '../utils/query'
	import type { Page } from '../utils/query'
	import type { QueryFunction } from '@tanstack/svelte-query'

	interface Song {
		id: number
		title: string
		artist: string
		album: string
		picture: number
		mime: string
	}

	const fetchSongs = (async ({ pageParam = 1 }) => {
		return fetch(`/api/songs?page=${pageParam}`)
		.then(res => res.json())
	}) satisfies QueryFunction

	const query = createInfiniteQuery<Page<Song[]>>({
		queryKey: ['songs'],
		queryFn: fetchSongs,
		getNextPageParam
	})
</script>

<Section>
	{#if $query.isSuccess}
		<div id="container">
			<div id="songs">
				{#each $query.data.pages as page}
					{#each page.data as { id, title, artist, picture, mime }}
						<SongSnippet {id} {title} {artist} {picture} {mime} />
					{/each}
				{/each}
			</div>
		</div>
		{#if $query.hasNextPage}
			<Button 
			 	--align-self="center" 
  				value="Load more..." 
    				on:click={() => $query.fetchNextPage()} 
    				disabled={$query.isFetchingNextPage}
    			/>
		{/if}
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
