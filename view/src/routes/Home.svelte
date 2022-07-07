<script lang="ts">
	import { onMount } from 'svelte'
	import Stat from '../components/Stat.svelte'

	let stats = {
		totalArtists: 0,
		totalSongs: 0,
		totalAlbums: 0
	}

	onMount(async () => {
		const res = await fetch('/api/stats')
		stats = await res.json()
	})
</script>

<section>
	<h1>Welcome to your music library</h1>
	<div id="stats">
		<Stat title="Total Songs" value={stats.totalSongs} />
		<Stat title="Total Artists" value={stats.totalArtists} />
		<Stat title="Total Albums" value={stats.totalAlbums} />
	</div>
</section>

<style>
	section {
		padding: 20px;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	h1 {
		color: #f9f9f9;
	}

	#stats {
		display: flex;
		gap: 10px;
	}
</style>
