<script lang="ts">
	import { Router, Route } from 'svelte-navigator'
	import { currentSong } from './components/player/audio'

	import Player from './components/player/Player.svelte'
	import Navbar from './components/Navbar.svelte'

	// Routes
	import AlbumsRoute from './routes/Albums.svelte'
	import AlbumRoute from './routes/Album.svelte'
	import ArtistsRoute from './routes/Artists.svelte'
	import ArtistRoute from './routes/Artist.svelte'
	import SongsRoute from './routes/Songs.svelte'
	import HomeRoute from './routes/Home.svelte'
	import PlaylistsRoute from './routes/Playlists.svelte'
</script>

<main>
	<Router>
		<div class="layout" class:noSongLayout={$currentSong === undefined}>
			<Navbar />
			<Route path="/" component={HomeRoute} />
			<Route path="/albums" component={AlbumsRoute} />
			<Route path="/album/:albumID" component={AlbumRoute} />
			<Route path="/artists" component={ArtistsRoute} />
			<Route path="/artist/:artistID" component={ArtistRoute} />
			<Route path="/songs" component={SongsRoute} />
			<Route path="/playlists" component={PlaylistsRoute} />
		</div>
		<Player />
	</Router>
</main>

<style>
	:global(body) {
		font-family: inter, sans-serif;
		--player-height: 70px;
		--main-color: #006e7f;
		--second-color: #0f0f0e;
		--third-color: #0c0c0c;
		--forth-color: #161616;
		--fifth-color: #1a1a1a;
	}

	:global(p) {
		margin: 0;
		padding: 0;
	}

	:global(:focus) {
		outline: none;
	}

	:global(a) {
		text-decoration: none;
	}

	main {
		height: 100vh;
		display: grid;
		grid-template-rows: 1fr var(--player-height);
		background-color: var(--second-color);
	}

	.layout {
		display: grid;
		grid-template-columns: 220px 1fr;
		height: calc(100vh - var(--player-height));
	}

	.noSongLayout {
		height: 100vh;
	}
</style>
