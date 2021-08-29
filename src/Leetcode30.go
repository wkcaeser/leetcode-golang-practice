package main

import "fmt"

func main() {
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}

func findSubstring(s string, words []string) []int {
	var rs []int

	if len(words) == 0 {
		return rs
	}

	wordMap := make(map[string]int)
	for _, word := range words {
		cnt, ok := wordMap[word]
		if ok {
			wordMap[word] = cnt + 1
		} else {
			wordMap[word] = 1
		}
	}

	wordLen := len(words[0])
	wordsLen := len(words) * wordLen

	strLen := len(s)
	if wordsLen > strLen {
		return rs
	}

	for i := 0; i < strLen-wordsLen+1; i++ {
		wordMapTemp := make(map[string]int)
		for j := 0; j < wordsLen; j += wordLen {
			subStr := s[i+j : i+j+wordLen]
			_, ok := wordMap[subStr]
			if ok {
				cnt, ok := wordMapTemp[subStr]
				if ok {
					wordMapTemp[subStr] = cnt + 1
				} else {
					wordMapTemp[subStr] = 1
				}
			} else {
				break
			}
		}
		if len(wordMapTemp) == len(wordMap) {
			match := true
			for k, v := range wordMapTemp {
				if wordMap[k] != v {
					match = false
					break
				}
			}
			if match {
				rs = append(rs, i)
			}
		}
	}

	return rs
}
