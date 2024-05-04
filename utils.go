package main

func areAllFalse(items []bool) bool {
	for _, item := range items {
		if item {
			return false
		}
	}
	return true
}
