<script lang="ts">
	import { Link } from 'svelte-navigator'
	import { createQuery } from '@tanstack/svelte-query'
	import Album from '../components/Album.svelte'
	import AlbumsWrapper from '../components/AlbumsWrapper.svelte'
	import Section from '../components/Section.svelte'

	interface AlbumPreview {
		id: number
		name: string
		artist: string
		picture: number
	}

	const query = createQuery<AlbumPreview[]>({
		queryKey: ['albums'],
		queryFn: () => fetch('/api/albums').then(res => res.json())
	})
</script>

<Section>
	<AlbumsWrapper>
		{#if $query.isSuccess}
			{#each $query.data as { id, name, artist, picture }}
				<Link to="/album/{id}">
					<Album {name} {artist} {picture} />
				</Link>
			{/each}
		{/if}
	</AlbumsWrapper>
</Section>
