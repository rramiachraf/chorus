import { writable, get } from 'svelte/store'
import { songs } from '../../routes/store'

export const currentSong = writable(undefined)
export const songPicture = writable(undefined)
export const songArtist = writable(undefined)
export const songTitle = writable(undefined)
export const songDuration = writable(0)
export const songCurrentTime = writable(0)
export const songPlaying = writable(false)
export const songMuted = writable(false)
export const songVolume = writable(1)
export const songRepeat = writable(false)
export const songShuffle = writable(false)

export const audioPlayer = new Audio()
audioPlayer.preload = 'metadata'

interface AudioEvent extends Event {
	target: HTMLAudioElement
}

audioPlayer.addEventListener('pause', () => {
	songPlaying.set(false)
})

audioPlayer.addEventListener('play', () => {
	songPlaying.set(true)
})

audioPlayer.addEventListener('loadeddata', async () => {
	try {
		const result = await Notification.requestPermission()
		let playing: boolean
		songPlaying.subscribe(value => {
			playing = value
		})()
		if (result === 'granted' && playing) {
			let title: string
			let artist: string
			let icon: string
			songTitle.subscribe(value => {
				title = value
			})()
			songArtist.subscribe(value => {
				artist = value
			})()
			songPicture.subscribe(value => {
				icon = document.location.origin + value
			})()

			if (title) {
				new Notification(artist, {
					body: title,
					tag: 'music',
					icon
				})
			}
		}
	} catch (e) {
		console.error(e)
	}
})

audioPlayer.addEventListener('timeupdate', (e: AudioEvent) => {
	if (e) {
		const currentTime = e.target.currentTime
		songCurrentTime.set(currentTime)
		localStorage.setItem('currentTime', String(currentTime))
	}
})

audioPlayer.addEventListener('durationchange', (e: AudioEvent) => {
	e && songDuration.set(e.target.duration)
})

audioPlayer.addEventListener('volumechange', (e: AudioEvent) => {
	if (e) {
		const volume = e.target.volume
		songVolume.set(volume)
		localStorage.setItem('currentVolume', String(volume))
	}
})

audioPlayer.addEventListener('ended', async () => {
	if (get(songShuffle)) {
		const res = await fetch('/api/random/song')
		const { id } = await res.json()
		playSong(id)
	}
})

export function togglePlay() {
	if (get(songPlaying)) {
		audioPlayer.pause()
	} else {
		audioPlayer.play()
	}
}

function getSongURL(songID: number) {
	return `/api/listen/${songID}`
}

currentSong.subscribe(async songID => {
	if (!songID) {
		return
	}

	const res = await fetch(`/api/song/${songID}`)
	const { picture, artist, title } = await res.json()

	songPicture.set(`/api/picture/${picture}`)
	songArtist.set(artist)
	songTitle.set(title)
})

export const playSong = (songID: number) => {
	audioPlayer.src = getSongURL(songID)
	audioPlayer.play()
	currentSong.set(songID)
	localStorage.setItem('currentSong', String(songID))
}

export const loadLastSong = () => {
	const songID = localStorage.getItem('currentSong')
	if (songID) {
		audioPlayer.src = getSongURL(Number(songID))
		audioPlayer.currentTime = Number(localStorage.getItem('currentTime'))
		audioPlayer.volume = Number(localStorage.getItem('currentVolume'))
		currentSong.set(songID)
	}
}

export const updateCurrentTime = (time: number) => {
	audioPlayer.currentTime = time
}

export const forward = () => {
	audioPlayer.currentTime += 5
}

export const backward = () => {
	audioPlayer.currentTime -= 5
}

export const toggleMute = () => {
	audioPlayer.muted = !audioPlayer.muted
	songMuted.set(audioPlayer.muted)
}

export const changeVolume = (newVolume: number) => {
	audioPlayer.volume = newVolume
}

export const toggleRepeat = () => {
	if (get(songShuffle)) {
		toggleShuffle()
	}

	audioPlayer.loop = !audioPlayer.loop
	songRepeat.update(loop => !loop)
}

export const toggleShuffle = () => {
	if (!get(songShuffle) && get(songRepeat)) {
		toggleRepeat()
	}
	songShuffle.update(shuffle => !shuffle)
}
