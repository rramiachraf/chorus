import { writable } from 'svelte/store'

export const albumScroll = writable(0)

export const albums = writable([])
export const artists = writable([])
export const songs = writable([])
