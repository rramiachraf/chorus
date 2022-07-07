<script lang="ts">
	import Icon from '@iconify/svelte'
	import RepeatOne from '@iconify/icons-lucide/repeat-1'
	import { songRepeat, toggleRepeat } from './audio'
</script>

<button on:click={toggleRepeat} class:active={$songRepeat}>
	<Icon icon={RepeatOne} />
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
