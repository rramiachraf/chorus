function formatTime(seconds: number): string {
	let date = new Date(null)
	date.setSeconds(isNaN(seconds) || seconds === Infinity ? 0 : seconds)
	return date.toISOString().substring(14, 19)
}

export default formatTime
