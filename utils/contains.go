package utils

func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if item == v {
			return true
		}
	}

	return false
}
