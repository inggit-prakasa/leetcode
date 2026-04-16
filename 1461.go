package main

func hasAllCodes(s string, k int) bool {
	mapCodes := make(map[string]bool)
	for i := 0; i <= len(s)-k; i++ {
		mapCodes[s[i:i+k]] = true
	}
		
	return len(mapCodes) == 1<<k
}
