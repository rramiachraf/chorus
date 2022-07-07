<script lang="ts">
	import Icon from '@iconify/svelte'
	import VolumeUp from '@iconify/icons-bi/volume-up-fill'
	import VolumeMute from '@iconify/icons-bi/volume-mute-fill'
	import { songMuted, toggleMute, songVolume, changeVolume } from './audio'

	let volumeWidth: number
	let dragContainer: HTMLDivElement
	$: volumeWidth = ($songVolume / 1) * 100

	function handleVolume({ clientX }) {
		const rect = dragContainer.getBoundingClientRect()
		const relativeMousePosition = clientX - rect.left
		const newVolume = relativeMousePosition / dragContainer.offsetWidth
		changeVolume(newVolume)
	}

	function makeVolumeFull() {
		changeVolume(1)
	}
</script>

<div class="volume">
	<div class="mute" on:click={toggleMute}>
		{#if $songMuted}
			<Icon icon={VolumeMute} />
		{:else}
			<Icon icon={VolumeUp} />
		{/if}
	</div>
	<div
		class="drag-container"
		on:click={handleVolume}
		on:drag={handleVolume}
		on:dblclick={makeVolumeFull}
		bind:this={dragContainer}
	>
		<div class="volume-level" style="width:{volumeWidth}%" />
	</div>
</div>

<style>
	.volume {
		color: white;
		display: flex;
		align-items: center;
		gap: 5px;
	}

	.drag-container {
		height: 6px;
		background-color: var(--fifth-color);
		width: 80px;
		border-radius: 100px;
		cursor: pointer;
	}

	.volume-level {
		background-color: var(--main-color);
		width: 0;
		height: 100%;
		border-radius: 100px;
		transition: 0.2s ease width;
	}

	.mute {
		display: flex;
		cursor: pointer;
	}
</style>
