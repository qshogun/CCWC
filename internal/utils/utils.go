package utils

func IsAnyTrue(items []bool) bool {
	for _, item := range items {
		if item {
			return true
		}
	}
	return false
}
