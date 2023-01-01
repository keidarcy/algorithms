package quicksort

func Sort(values []int) {
	quickSort(values, 0, len(values)-1)
}

func quickSort(values []int, start, end int) {
	if start >= end {
		return
	}
	// partition
	boundary := partition(values, start, end)

	// sort left
	quickSort(values, start, boundary-1)

	// sort right
	quickSort(values, boundary, end)
}

func partition(values []int, start, end int) (boundary int) {
	pivot := values[end]
	boundary = start - 1
	for i := start; i <= end; i++ {
		current := values[i]
		if current <= pivot {
			boundary++
			values[boundary], values[i] = values[i], values[boundary]
		}
	}
	return boundary
}
