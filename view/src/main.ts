import 'normalize.css'
import '@fontsource-variable/inter'
import App from './App.svelte'
import {
	forward,
	backward,
	togglePlay,
	toggleRepeat,
	toggleShuffle,
	toggleMute,
	increaseVolume,
	decreaseVolume,
	updateCurrentTime,
    playNext,
    playPrevious
} from './components/player/audio'
import { isSearching } from './store'

const app = new App({
	target: document.body
})

let searching = false
isSearching.subscribe(v => searching = v)

function handleKeyboard(e: KeyboardEvent) {
	if (searching){
		return
	}

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
			increaseVolume()
			break
		case 'j':
			decreaseVolume()
			break
		case 'ArrowRight':
			forward()
			break
		case 'ArrowLeft':
			backward()
			break
		case 'ArrowUp':
			increaseVolume()
			break
		case 'ArrowDown':
			decreaseVolume()
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
		case 'x':
			updateCurrentTime(0)
			break
		case 'b':
			playNext()
			break
		case 'z':
			playPrevious()
			break
	}
}

window.addEventListener('keydown', handleKeyboard)

export default app
