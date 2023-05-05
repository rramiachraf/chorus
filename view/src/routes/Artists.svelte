<script lang="ts">
	import { onMount } from 'svelte'
	import { Link } from 'svelte-navigator'
	import Artist from '../components/Artist.svelte'
	import Section from '../components/Section.svelte'
	import { artists } from './store'

	onMount(async () => {
		if ($artists.length === 0) {
			const res = await fetch('/api/artists')
			artists.set(await res.json())
		}
	})
</script>

<Section>
	<div>
		{#each $artists as { name, id }}
			<Link to="/artist/{id}">
				<Artist {name} picture="" />
			</Link>
		{/each}
	</div>
</Section>

<style>
	div {
		display: grid;
		grid-template-columns: repeat(6, 150px);
		justify-content: space-between;
		gap: 10px;
	}
</style>
