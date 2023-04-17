<script lang="ts">
	import {
		songCurrentTime,
		songDuration,
		updateCurrentTime
	} from './audio'
	import formatTime from '../../utils/formatTime'

	let progressWidth = 0
	let progressContainer: HTMLDivElement
	$: progressWidth = ($songCurrentTime / $songDuration) * 100

	function handleProgress({ clientX }) {
		const rect = progressContainer.getBoundingClientRect()
		const relativeMousePosition = clientX - rect.left
		const postionPercentage =
			(relativeMousePosition /
				progressContainer.offsetWidth) *
			100
		const newCurrentTime = (postionPercentage * $songDuration) / 100
		updateCurrentTime(newCurrentTime)
	}
</script>

<section>
	<p>{formatTime($songCurrentTime)}</p>
	<div id="bg" on:click={handleProgress} bind:this={progressContainer}>
		<div id="progress" style="width: {progressWidth.toFixed(3)}%" />
	</div>
	<p>{formatTime($songDuration)}</p>
</section>

<style>
	section {
		display: grid;
		justify-content: center;
		align-items: center;
		gap: 5px;
		grid-template-columns: 40px 1fr 40px;
	}

	p {
		font-size: 12px;
		color: #bbb;
	}

	p:first-child {
		justify-self: flex-start;
	}

	p:last-child {
		justify-self: flex-end;
	}

	#bg {
		background-color: var(--fifth-color);
		height: 6px;
		width: 100%;
		border-radius: 100px;
		cursor: pointer;
		transition: 0.1s ease height;
	}

	#progress {
		height: 100%;
		width: 0%;
		background: var(--main-color);
		border-radius: 100px;
		transition: 0.2s ease width;
		position: relative;
	}

	#progress::after {
		content: '';
		width: 10px;
		height: 10px;
		position: absolute;
		top: 50%;
		right: -5px;
		background: var(--third-color);
		border-radius: 50%;
		border: 2px solid var(--main-color);
		transform: translateY(-50%);
	}
</style>
