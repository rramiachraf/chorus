<script>
	import Icon from '@iconify/svelte'
	import ShuffleIcon from '@iconify/icons-lucide/shuffle'
	import { songShuffle, toggleShuffle } from './audio'
</script>

<button on:click={toggleShuffle} class:active={$songShuffle}>
	<Icon icon={ShuffleIcon} />
</button>

<style>
	button {
		background: none;
		border: none;
		cursor: pointer;
		display: flex;
		color: #444;
		transition: 0.2s ease-out color;
	}

	.active {
		color: white;
	}
</style>
