package main

func isAnyTrue(items []bool) bool {
	for _, item := range items {
		if item {
			return true
		}
	}
	return false
}
