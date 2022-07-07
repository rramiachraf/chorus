<script lang="ts">
	import { songArtist, songTitle, songPicture } from './audio'

	function trim(str: string) {
		if (!$songTitle) {
			return ''
		}

		return str.length >= 30 ? str.substring(0, 35) + '...' : str
	}
</script>

<section>
	<figure style="background-image: url('{$songPicture}')" />
	<div id="about">
		<p id="artist">{$songArtist}</p>
		<p id="title">{trim($songTitle)}</p>
	</div>
</section>

<style>
	section {
		display: flex;
		align-items: center;
		gap: 10px;
	}

	figure {
		height: 40px;
		width: 40px;
		border-radius: 5px;
		background-size: cover;
		background-color: var(--second-color);
		margin: 0;
	}

	#about {
		display: flex;
		flex-direction: column;
		gap: 2px;
	}

	#title {
		font-weight: 600;
		color: #ddd;
		font-size: 14px;
	}

	#artist {
		color: #999;
		font-size: 12px;
	}
</style>
