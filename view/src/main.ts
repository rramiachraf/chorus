import 'normalize.css'
import '@fontsource/inter/variable.css'
import App from './App.svelte'
import {
	forward,
	backward,
	changeVolume,
	songVolume,
	togglePlay,
	toggleRepeat,
	toggleShuffle,
	toggleMute
} from './components/player/audio'

const app = new App({
	target: document.body
})

function getSongVolume() {
	let volume: number
	songVolume.subscribe(v => (volume = v))()
	return volume
}

function handleKeyboard(e: KeyboardEvent) {
	switch (e.key) {
		case ' ':
			togglePlay()
			break
		case 'l':
			forward()
			break
		case 'h':
			backward()
			break
		case 'k':
			const volumeUp = getSongVolume() + 0.05
			if (volumeUp < 1) {
				changeVolume(volumeUp)
			} else {
				changeVolume(1)
			}
			break
		case 'j':
			const volumeDown = getSongVolume() - 0.05
			if (volumeDown < 0) {
				changeVolume(0)
			} else {
				changeVolume(volumeDown)
			}
			break
		case 'm':
			toggleMute()
			break
		case 'r':
			toggleRepeat()
			break
		case 's':
			toggleShuffle()
			break
	}
}

window.addEventListener('keydown', handleKeyboard)

export default app
