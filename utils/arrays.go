package utils

func FindFirstIndex[T comparable](slice []T, evalFunc func(T) bool) (int, bool) {
	for i := 0; i < len(slice); i++ {
		if evalFunc(slice[i]) {
			return i, true
		}
	}

	return -1, false
}

func FindLastIndex[T comparable](slice []T, evalFunc func(T) bool) (int, bool) {
	for i := len(slice) - 1; i >= 0; i-- {
		if evalFunc(slice[i]) {
			return i, true
		}
	}

	return -1, false
}
