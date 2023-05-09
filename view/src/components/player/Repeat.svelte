<script lang="ts">
	import Icon from '@iconify/svelte'
	import RepeatOne from '@iconify/icons-lucide/repeat-1'
	import Repeat from '@iconify/icons-lucide/repeat'
	import { songRepeat, toggleRepeat, playbackMode } from './audio'
</script>

<button
	on:click={toggleRepeat}
	class:active={$songRepeat !== playbackMode.NONE}
>
	{#if $songRepeat === playbackMode.REPEAT_ONE}
		<Icon icon={RepeatOne} />
	{:else}
		<Icon icon={Repeat} />
	{/if}
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
