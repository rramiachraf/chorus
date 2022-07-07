<script lang="ts">
	import { onMount } from 'svelte'
	import Album from '../components/Album.svelte'
	import AlbumsWrapper from '../components/AlbumsWrapper.svelte'

	export let artistID: number

	let artistName: string
	let albums = []

	onMount(async () => {
		const res = await fetch(`/api/artist/${artistID}`)
		const json = await res.json()
		artistName = json.name
		json.albums && (albums = json.albums)
	})
</script>

<section>
	<h1>{artistName}</h1>
	<AlbumsWrapper>
		{#each albums as { id, name, picture }}
			<Album {id} {name} picture="/api/picture/{picture}" />
		{/each}
	</AlbumsWrapper>
</section>

<style>
	h1 {
		color: white;
		text-align: center;
		margin: 0;
	}

	section {
		overflow: auto;
		padding: 10px;
		display: flex;
		flex-direction: column;
		gap: 20px;
	}
</style>
