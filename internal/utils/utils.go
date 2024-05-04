package utils

func IsAnyTrue(items map[string]bool) bool {
	for _, item := range items {
		if item {
			return true
		}
	}
	return false
}
