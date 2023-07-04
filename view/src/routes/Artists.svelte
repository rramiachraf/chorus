<script lang="ts">
	import { createQuery } from '@tanstack/svelte-query'
	import { Link } from 'svelte-navigator'
	import Artist from '../components/Artist.svelte'
	import Section from '../components/Section.svelte'

	interface Artist {
		id: number
		name: string
	}

	const query = createQuery<Artist[]>({
		queryKey: ['artists'],
		queryFn: () => fetch('/api/artists').then(res => res.json())
	})
</script>

<Section>
	{#if $query.isSuccess}
		<div>
			{#each $query.data as { name, id }}
				<Link to="/artist/{id}">
					<Artist {name} picture="" />
				</Link>
			{/each}
		</div>
	{/if}
</Section>

<style>
	div {
		display: grid;
		grid-template-columns: repeat(6, 150px);
		justify-content: space-between;
		gap: 10px;
	}
</style>
