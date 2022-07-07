<script lang="ts">
	import { onMount } from 'svelte'
	import { Link } from 'svelte-navigator'
	import Artist from '../components/Artist.svelte'
	import { artists } from './store'

	onMount(async () => {
		if ($artists.length === 0) {
			const res = await fetch('/api/artists')
			artists.set(await res.json())
		}
	})
</script>

<section>
	{#each $artists as { name, id }}
		<Link to="/artist/{id}"><Artist {name} picture="" /></Link>
	{/each}
</section>

<style>
	section {
		padding: 20px;
		display: grid;
		grid-template-columns: repeat(6, 150px);
		gap: 10px;
		overflow-y: auto;
		justify-content: space-between;
	}
</style>
