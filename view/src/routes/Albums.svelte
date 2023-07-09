<script lang="ts">
	import { Link } from 'svelte-navigator'
	import { createInfiniteQuery } from '@tanstack/svelte-query'
	import Album from '../components/Album.svelte'
	import AlbumsWrapper from '../components/AlbumsWrapper.svelte'
	import Section from '../components/InfiniteSection.svelte'
	import Button from '../components/Button.svelte'
	import type { QueryFunction } from '@tanstack/svelte-query'
	import { getNextPageParam } from '../utils/query'
	import type { Page } from '../utils/query'

	interface AlbumPreview {
		id: number
		name: string
		artist: string
		picture: number
	}

	const fetchAlbums = (async ({ pageParam = 1 }) => {
		return fetch(`/api/albums?page=${pageParam}`)
		.then(res => res.json())
	}) satisfies QueryFunction

	const query = createInfiniteQuery<Page<AlbumPreview[]>>({
		queryKey: ['albums'],
		queryFn: fetchAlbums, 
		getNextPageParam
	})
</script>

<Section>
	{#if $query.isSuccess}
		<AlbumsWrapper>
			{#each $query.data.pages as page}
				{#each page.data as {id, name, artist, picture}}
					<Link to="/album/{id}">
						<Album {name} {artist} {picture} />
					</Link>
				{/each}
			{/each}
		</AlbumsWrapper>
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
