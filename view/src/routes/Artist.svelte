<script lang="ts">
	import { onMount } from 'svelte'
	import { useNavigate } from 'svelte-navigator'
	import Album from '../components/Album.svelte'
	import AlbumsWrapper from '../components/AlbumsWrapper.svelte'
	import { playSong } from '../components/player/audio'

	const navigate = useNavigate()

	export let artistID: number

	let artistName: string
	let albums = []
	let singles = []

	onMount(async () => {
		const res = await fetch(`/api/artist/${artistID}`)
		const json = await res.json()
		artistName = json.name
		json.albums && (albums = json.albums)
		json.singles && (singles = json.singles)
	})
</script>

<section>
	<div class="info">
		<h1>{artistName}</h1>
	</div>
	<div class="content">
		{#if albums.length > 0}<h1 class="title">Albums</h1>{/if}
		<AlbumsWrapper>
			{#each albums as { id, name, picture }}
				<Album
					{id}
					{name}
					{picture}
					on:click={() => navigate(`/album/${id}`)}
				/>
			{/each}
		</AlbumsWrapper>
		{#if singles.length > 0}<h1 class="title">Singles</h1>{/if}
		<AlbumsWrapper>
			{#each singles as { id, title, picture }}
				<Album {id} name={title} {picture} on:click={() => playSong(id)} />
			{/each}
		</AlbumsWrapper>
	</div>
</section>

<style>
	.info {
		background: url(/bg-pattern.jpg);
		display: flex;
		justify-content: center;
		align-items: center;
		border-bottom: 1px solid var(--fifth-color);
	}

	.info h1 {
		color: white;
		font-size: 30px;
	}

	.title {
		color: #eee;
		font-size: 16px;
		font-weight: 600;
		text-transform: uppercase;
	}

	section {
		overflow: auto;
		display: grid;
		grid-template-rows: 40vh 1fr;
	}

	.content {
		padding: 20px;
		display: flex;
		flex-direction: column;
		gap: 20px;
	}
</style>
