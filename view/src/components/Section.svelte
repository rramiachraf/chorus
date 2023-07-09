<script lang="ts">
	import { onMount } from 'svelte'
	import { useLocation } from 'svelte-navigator'
	import { scrollPosition } from '../store'

	const location = useLocation()
	let section: HTMLElement
	let style: string

	onMount(() => {
		setTimeout(() => {
			section.scrollTop = $scrollPosition.get($location.pathname) || 0
		}, 0)
	})

	function handleScroll(e: UIEvent) {
		const position = (e.target as HTMLElement).scrollTop
		scrollPosition.update(s => s.set($location.pathname, position))
	}
</script>

<section on:scroll={handleScroll} bind:this={section}>
	<slot />
</section>

<style>
	section {
		overflow-y: auto;
		padding: 20px;
		scrollbar-color: var(--main-color) transparent;
		scrollbar-width: auto;
		scroll-behavior: initial;
		display: var(--display);
		flex-direction: var(--direction);
		gap: var(--gap);
	}
</style>
