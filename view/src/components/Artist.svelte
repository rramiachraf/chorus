<script lang="ts">
	import Icon from '@iconify/svelte'
	import MusicNote from '@iconify/icons-bi/music-note'

	export let name: string
	export let picture: string
</script>

<article>
	<figure style="background-image: url('{picture}');">
		{#if !picture} <Icon icon={MusicNote} width="100" /> {/if}
	</figure>
	<h2>{name}</h2>
</article>

<style>
	article {
		display: flex;
		flex-direction: column;
		align-items: center;
		cursor: pointer;
		transition: 0.1s ease opacity;
	}

	article:hover {
		opacity: 0.7;
	}

	figure {
		width: 150px;
		height: 150px;
		background-color: var(--fifth-color);
		background-size: cover;
		border-radius: 50%;
		display: flex;
		align-items: center;
		justify-content: center;
		color: #777;
	}

	h2 {
		margin: 0;
		padding: 0;
		font-size: 15px;
		font-weight: 500;
		color: #eee;
		text-align: center;
	}
</style>
