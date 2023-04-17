<script lang="ts">
	import { onMount } from 'svelte'
	import numbro from 'numbro'
	import Stat from '../components/Stat.svelte'

	let stats: Record<string, number | string> = {
		totalArtists: 0,
		totalSongs: 0,
		totalAlbums: 0
	}

	const opts: numbro.Format = { thousandSeparated: true }

	$: stats.totalSongs = numbro(stats.totalSongs).format(opts)
	$: stats.totalAlbums = numbro(stats.totalAlbums).format(opts)
	$: stats.totalArtists = numbro(stats.totalArtists).format(opts)

	onMount(async () => {
		try {
			const res = await fetch('/api/stats')
			stats = await res.json()
		} catch (e) {}
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
