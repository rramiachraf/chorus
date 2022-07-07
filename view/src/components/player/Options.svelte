<script lang="ts">
	import Volume from './Volume.svelte'
	import Repeat from './Repeat.svelte'
	import Shuffle from './Shuffle.svelte'
</script>

<section>
	<Shuffle />
	<Repeat />
	<Volume />
</section>

<style>
	section {
		display: flex;
		align-items: center;
		justify-content: flex-end;
		gap: 8px;
	}
</style>
