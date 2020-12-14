package main

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	rs := -1
	for i := 0; i < len(haystack); i++ {
		if len(haystack)-i < len(needle) {
			break
		}
		if haystack[i] == needle[0] {
			isEqual := true
			for j := 1; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					isEqual = false
					break
				}
			}
			if isEqual {
				rs = i
			}
		}
		if rs != -1 {
			break
		}
	}

	return rs
}
