import { writable, get } from 'svelte/store'
import { tracksList } from '../../routes/store'

export enum playbackMode {
	NONE,
	REPEAT_ONE,
	REPEAT_ALL
}

export const currentSong = writable(undefined)
export const songPicture = writable(undefined)
export const songArtist = writable(undefined)
export const songTitle = writable(undefined)
export const songDuration = writable(0)
export const songCurrentTime = writable(0)
export const songPlaying = writable(false)
export const songMuted = writable(false)
export const songVolume = writable(1)
export const songRepeat = writable(playbackMode.NONE)
export const songShuffle = writable(false)
export const songError = writable(false)

export const audioPlayer = new Audio()
audioPlayer.preload = 'metadata'

interface AudioEvent extends Event {
	target: HTMLAudioElement
}

export function getNextTrack(currentTrack: number) {
	const t = get(tracksList)
	if (t.length === 0) {
		return currentTrack
	}
	const i = t.indexOf(currentTrack)
	const next = i + 1
	if (next === t.length) {
		return t[0]
	}

	return t[next]
}

export function getPreviousTrack(currentTrack: number) {
	const t = get(tracksList)
	if (t.length === 0) {
		return currentTrack
	}
	const i = t.indexOf(currentTrack)
	const prev = i - 1
	if (prev < 0) {
		return t[0]
	}

	return t[prev]
}


audioPlayer.addEventListener('error', () => {
	songError.set(true)
})

audioPlayer.addEventListener('pause', () => {
	songPlaying.set(false)
})

audioPlayer.addEventListener('play', () => {
	songPlaying.set(true)
	songError.set(false)
})

audioPlayer.addEventListener('loadeddata', async () => {
	try {
		const result = await Notification.requestPermission()
		if (result === 'granted' && get(songPlaying)) {
			const artist = get(songArtist)
			const title = get(songTitle)
			const icon = document.location.origin + get(songPicture)

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

	if (get(songRepeat) === playbackMode.REPEAT_ALL && get(tracksList).length > 0){
	       playSong(getNextTrack(get(currentSong)))
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
		audioPlayer.currentTime = Number(
			localStorage.getItem('currentTime')
		)
		audioPlayer.volume = Number(
			localStorage.getItem('currentVolume')
		)
		currentSong.set(Number(songID))
		const t = localStorage.getItem('tracksList')
		if (t) {
			tracksList.set(JSON.parse(t))
		}
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

	switch (get(songRepeat)) {
		case playbackMode.NONE:
			songRepeat.set(playbackMode.REPEAT_ONE)
			audioPlayer.loop = true
			break
		case playbackMode.REPEAT_ONE:
			songRepeat.set(playbackMode.REPEAT_ALL)
			audioPlayer.loop = false
			break
		default:
			songRepeat.set(playbackMode.NONE)
			audioPlayer.loop = false
	}
}

export const toggleShuffle = () => {
	if (!get(songShuffle) && get(songRepeat)) {
		songRepeat.set(playbackMode.NONE)
		audioPlayer.loop = false
	}
	songShuffle.update(shuffle => !shuffle)
}

export const increaseVolume = () => {
	const volumeUp = get(songVolume) + 0.05
	if (volumeUp < 1) {
		changeVolume(volumeUp)
	} else {
		changeVolume(1)
	}
}

export const decreaseVolume = () => {
	const volumeDown = get(songVolume) - 0.05
	if (volumeDown < 0) {
		changeVolume(0)
	} else {
		changeVolume(volumeDown)
	}
}

export function playNext() {
		const next = getNextTrack(get(currentSong))
		playSong(next)
}

export function playPrevious() {
		const next = getPreviousTrack(get(currentSong))
		playSong(next)
}

