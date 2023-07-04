<script lang="ts">
	import { createQuery } from '@tanstack/svelte-query'
	import Stat from '../components/Stat.svelte'

	interface Stats {
		totalArtists: number
		totalSongs: number
		totalAlbums: number
	}

	const query = createQuery<Stats>({
		queryKey: ['stats'],
		queryFn: () => fetch('/api/stats').then(res => res.json())
	})
</script>

<section>
	<h1>Welcome to your music library</h1>
	<div id="stats">
		{#if $query.isSuccess}
			<Stat title="Total Songs" value={$query.data.totalSongs} />
			<Stat title="Total Artists" value={$query.data.totalArtists} />
			<Stat title="Total Albums" value={$query.data.totalAlbums} />
		{/if}
	</div>
</section>

<style>
	section {
		padding: 20px;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		gap: 15px;
	}

	h1 {
		color: #f9f9f9;
	}

	#stats {
		display: flex;
		gap: 10px;
	}
</style>
