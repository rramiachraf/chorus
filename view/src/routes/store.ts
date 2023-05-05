import { writable } from 'svelte/store'

export const albums = writable([])
export const artists = writable([])
export const songs = writable([])

export const scrollPosition = writable(new Map<string, number>())
