<div>
	<slot />
</div>

<style>
	div {
		display: grid;
		grid-template-columns: repeat(5, 200px);
		justify-content: center;
		gap: 20px;
	}
</style>
