package main

import "fmt"

func symbolsToIntMap() map[string]int {
	return map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000}
}

func romanToInt(s string) int {
	ans := 0;
	intMap := symbolsToIntMap()
	length := len(s) - 1
	finish := false
	for i := 0; i < length; i++ {
		value, ok := intMap[s[i:i+2]]
		if ok {
			ans += value
			i++
			if i == length {
				finish = true
			}
			continue
		}
		ans += intMap[s[i:i+1]]
	}
	if !finish {
		ans += intMap[s[length:length+1]]
	}
	return ans
}

func main() {
	s := "CM"
	fmt.Println(romanToInt(s))
}
