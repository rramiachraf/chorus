<script lang="ts">
	export let name: string
	export let picture: string
	export let artist!: string
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div id="album" title="{name}" on:click>
	<div id="picture" style="background-image: url('/api/picture/{picture}')" />
	{#if artist}<p id="artist">{artist}</p>{/if}
	<p id="name">{name}</p>
</div>

<style>
	#album {
		display: flex;
		flex-direction: column;
		gap: 5px;
		width: 100%;
		cursor: pointer;
		border-radius: 10px;
		transition: 0.3s ease background-color;
		padding-bottom: 10px;
		border: 1px solid var(--fifth-color);
	}

	#album:hover {
		background-color: var(--fifth-color);
		opacity: 0.9;
	}

	p {
		margin: 0;
		padding: 0;
		text-align: center;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
		padding: 0 10px;
	}

	#picture {
		height: 200px;
		background-size: cover;
		border-radius: 10px;
	}

	#name {
		font-size: 15px;
		color: #ddd;
		font-weight: 600;
	}

	#artist {
		color: #bbb;
		font-size: 13px;
	}
</style>
