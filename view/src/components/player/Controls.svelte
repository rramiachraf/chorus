<script lang="ts">
	import Icon from '@iconify/svelte'
	import Progress from './Progress.svelte'
	import Play from '@iconify/icons-bi/play-fill'
	import Pause from '@iconify/icons-bi/pause-fill'
	import Backward from '@iconify/icons-bi/skip-backward-fill'
	import Forward from '@iconify/icons-bi/skip-forward-fill'
	import {
		songPlaying,
		togglePlay,
		playPrevious,
		playNext
	} from './audio'
</script>

<section>
	<div>
		<button id="skip" on:click={playPrevious}>
			<Icon icon={Backward} />
		</button>
		<button id="toggleplay" on:click={togglePlay}>
			{#if $songPlaying}
				<Icon icon={Pause} />
			{:else}
				<Icon icon={Play} />
			{/if}
		</button>
		<button id="skip" on:click={playNext}>
			<Icon icon={Forward} />
		</button>
	</div>
	<Progress />
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		gap: 2px;
	}

	div {
		display: flex;
		justify-content: center;
		align-items: center;
	}

	button {
		display: flex;
		align-items: center;
		justify-content: center;
		cursor: pointer;
	}

	#toggleplay {
		background: none;
		border: none;
		color: #fff;
		font-size: 24px;
	}

	#skip {
		background: none;
		color: #bbb;
		border: none;
		font-size: 16px;
	}
</style>
