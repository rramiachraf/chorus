<script lang="ts">
	import { Link, useLocation, navigate } from 'svelte-navigator'
	import Icon from '@iconify/svelte'
	import Disc from '@iconify/icons-bi/disc'
	import People from '@iconify/icons-bi/people-fill'
	import Note from '@iconify/icons-bi/music-note'
	import List from '@iconify/icons-bi/music-note-list'

	const links = [
		{ display: 'Albums', to: '/albums', icon: Disc },
		{ display: 'Artists', to: '/artists', icon: People },
		{ display: 'Songs', to: '/songs', icon: Note },
		{ display: 'Playlists', to: '/playlists', icon: List }
	]

	const location = useLocation()
	const currentLink = 'border-bottom: 2px solid var(--main-color)'

	function goHome() {
		navigate('/')
	}
</script>

<nav>
	<h1 on:click={goHome}>chorus</h1>
	{#each links as { display, to, icon }}
		<Link {to} style={$location.pathname === to && currentLink}
			><Icon {icon} />{display}</Link
		>
	{/each}
	<small>v0.1.0</small>
</nav>

<style>
	nav {
		background-color: var(--third-color);
		z-index: 10;
		display: flex;
		flex-direction: column;
		gap: 5px;
		padding: 10px 20px;
		border-right: 1px solid var(--fifth-color);
	}

	h1 {
		text-align: center;
		font-size: 20px;
		color: #eee;
		text-transform: uppercase;
		cursor: pointer;
		padding: 20px;
	}

	:global(nav > a) {
		color: #ccc;
		padding: 10px;
		background: var(--forth-color);
		border-radius: 5px;
		font-size: 14px;
		font-weight: 500;
		display: flex;
		align-items: center;
		gap: 10px;
		transition: 0.3s ease background-color;
		border: 2px solid transparent;
	}

	:global(nav > a:hover) {
		background-color: var(--main-color);
	}

	small {
		color: #bbb;
		text-align: center;
		justify-content: flex-end;
		display: flex;
		flex-direction: column;
		flex-grow: 1;
	}
</style>
