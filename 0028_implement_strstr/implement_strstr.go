package _028_implement_strstr

func strStr(haystack string, needle string) int {
	lenNeedle, lenHaystack := len(needle), len(haystack)
	if lenNeedle == 0 {
		return 0
	} else if lenNeedle > lenHaystack {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		for j := 0; i+j < lenHaystack && haystack[i + j] == needle[j]; j++ {
			if j + 1 == lenNeedle {
				return i
			}
		}
	}
	return -1
}
