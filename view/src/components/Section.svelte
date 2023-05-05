<script lang="ts">
	import { onMount } from 'svelte'
	import { useLocation } from 'svelte-navigator'
	import { scrollPosition } from '../routes/store'

	const location = useLocation()
	let section: HTMLElement

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
		scrollbar-color: var(--main-color) transparent;
		padding: 20px;
		scroll-behavior: initial;
	}
</style>
