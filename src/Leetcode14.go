package main

import "fmt"

func minLen(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func longestCommonPrefix(strs []string) string {
	ans := ""
	strsLen := len(strs)
	if strsLen == 0{
		return ans
	}
	minLength := 9999999
	for i := 0; i < strsLen; i ++ {
		minLength = minLen(minLength, len(strs[i]))
	}
	isContinue := true
	for i := 0; i < minLength; i++ {
		if isContinue {
			temp := strs[0][i:i+1]
			for j:=1; j<strsLen; j++ {
				if strs[j][i:i+1] != temp {
					isContinue = false
					break;
				}
			}
			if isContinue {
				ans += temp
			}
		}
	}
	return ans
}

func main()  {
	strs := []string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(strs))
}
