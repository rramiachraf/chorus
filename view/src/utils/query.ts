export interface Page<T> {
	data: T
	next: number
}

export function getNextPageParam<Data>(lastPage: Page<Data>) {
	if (lastPage.next) {
		return lastPage.next
	}

	return undefined
}
