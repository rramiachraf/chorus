<script lang="ts">
	export let title: string
	export let value: number | string
</script>

<div>
	<p>{title}</p>
	<h1>{value}</h1>
</div>

<style>
	div {
		background-color: var(--third-color);
		padding: 20px;
		border-radius: 5px;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		gap: 5px;
		border: 1px solid var(--forth-color);
	}

	h1 {
		font-size: 20px;
		color: #f9f9f9;
		margin: 0;
	}

	p {
		font-size: 12px;
		color: #ddd;
		text-transform: uppercase;
	}
</style>
