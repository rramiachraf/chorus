<script lang="ts">
	import { createInfiniteQuery } from '@tanstack/svelte-query'
	import { Link } from 'svelte-navigator'
	import Artist from '../components/Artist.svelte'
	import Section from '../components/InfiniteSection.svelte'
	import Button from '../components/Button.svelte'
	import { getNextPageParam } from '../utils/query'
	import type { QueryFunction } from '@tanstack/svelte-query'
	import type { Page } from '../utils/query'

	interface Artist {
		id: number
		name: string
	}

	const fetchArtists = (async ({ pageParam = 1 }) => {
		return fetch(`/api/artists?page=${pageParam}`)
		.then(res => res.json())
	}) satisfies QueryFunction

	const query = createInfiniteQuery<Page<Artist[]>>({
		queryKey: ['artists'],
		queryFn: fetchArtists,
		getNextPageParam
	})
</script>

<Section>
	{#if $query.isSuccess}
		<div>
			{#each $query.data.pages as page}
				{#each page.data as {id, name}}
					<Link to="/artist/{id}">
						<Artist {name} picture="" />
					</Link>
				{/each}
			{/each}
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
	div {
		display: grid;
		grid-template-columns: repeat(6, 150px);
		justify-content: space-between;
		gap: 10px;
	}
</style>
