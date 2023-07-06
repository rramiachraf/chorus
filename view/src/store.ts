import { writable } from 'svelte/store'

export const songs = writable([])
export const tracksList = writable(new Array<number>)

export const scrollPosition = writable(new Map<string, number>())
export const isSearching = writable(false)
