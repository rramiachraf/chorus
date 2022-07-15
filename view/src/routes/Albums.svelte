<script lang="ts">
	import { onMount } from 'svelte'
	import {useNavigate} from 'svelte-navigator'
	import Album from '../components/Album.svelte'
	import AlbumsWrapper from '../components/AlbumsWrapper.svelte'
	import { albums } from './store'
	
	const navigate = useNavigate()

	onMount(async () => {
		if ($albums.length === 0) {
			const res = await fetch('/api/albums')
			const json = await res.json()
			albums.set(json)
		}
	})
</script>

<section>
	<AlbumsWrapper>
		{#each $albums as { id, name, artist, picture }}
			<Album {name} {artist} {picture} on:click={() => navigate(`/album/${id}`)} />
		{/each}
	</AlbumsWrapper>
</section>

<style>
	section {
		overflow-y: auto;
		scrollbar-color: var(--main-color) transparent;
		padding: 20px;
	}
</style>
