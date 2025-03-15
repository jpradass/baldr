package utils

func Map[T any, Z any](slice []T, fn func(T) Z) []Z {
	result := make([]Z, len(slice))

	if len(slice) == 0 {
		return result
	}

	for idx, value := range slice {
		result[idx] = fn(value)
	}

	return result
}
