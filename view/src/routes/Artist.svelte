<script lang="ts">
	import { createQuery } from '@tanstack/svelte-query'
	import { useNavigate } from 'svelte-navigator'
	import Album from '../components/Album.svelte'
	import AlbumsWrapper from '../components/AlbumsWrapper.svelte'
	import { playSong } from '../components/player/audio'

	interface Artist {
		name: string
		albums: {
			id: number
			name: string
			picture: number
		}[]
		singles: {
			id: number
			title: string
			picture: number
		}[]
	}

	const navigate = useNavigate()

	export let artistID: number

	const query = createQuery<Artist>({
		queryKey: ['artist', artistID],
		queryFn: () => fetch(`/api/artist/${artistID}`).then(res => res.json())
	})
</script>

<section>
	{#if $query.isSuccess}
		<div class="info">
			<h1>{$query.data.name}</h1>
		</div>
		<div class="content">
			{#if $query.data.albums}
				<h1 class="title">Albums</h1>
				<AlbumsWrapper>
					{#each $query.data.albums as { id, name, picture }}
						<Album
							{name}
							{picture}
							on:click={() => navigate(`/album/${id}`)}
						/>
					{/each}
				</AlbumsWrapper>
			{/if}
			{#if $query.data.singles}
				<h1 class="title">Singles</h1>
				<AlbumsWrapper>
					{#each $query.data.singles as { id, title, picture }}
						<Album
							name={title}
							{picture}
							on:click={() => playSong(id)}
						/>
					{/each}
				</AlbumsWrapper>
			{/if}
		</div>
	{/if}
</section>

<style>
	.info {
		background: url('/bg-pattern.jpg');
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
