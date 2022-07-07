<script lang="ts">
	export let title: string
	export let value: any
</script>

<div>
	<h1>{value}</h1>
	<p>{title}</p>
</div>

<style>
	div {
		background-color: var(--third-color);
		padding: 20px;
		border-radius: 3px;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		gap: 5px;
		border: 1px solid #202020;
	}

	h1 {
		font-size: 20px;
		color: #f9f9f9;
		margin: 0;
	}

	p {
		font-size: 16px;
		color: #ddd;
	}
</style>
