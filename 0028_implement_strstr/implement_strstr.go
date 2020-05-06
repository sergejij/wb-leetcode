package main

func strStr(haystack string, needle string) int {
	lenNeedle, lenHaystack := len(needle), len(haystack)
	if lenNeedle == 0 {
		return 0
	} else if lenNeedle > lenHaystack {
		return -1
	}
	for i := 0; i < lenHaystack; i++ {
		if i + lenNeedle <= lenHaystack && haystack[i: i + lenNeedle] == needle {
			return i
		}
	}
	return -1
}
