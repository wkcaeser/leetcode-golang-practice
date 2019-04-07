package main

import "fmt"

func numMap() map[int]int {
	return map[int]int{
		1:  1,
		2:  4,
		3:  5,
		4:  9,
		5:  10,
		6:  40,
		7:  50,
		8:  90,
		9:  100,
		10: 400,
		11: 500,
		12: 900,
		13: 1000}
}

func symbolsMap() map[int]string {
	return map[int]string{
		1:  "I",
		2:  "IV",
		3:  "V",
		4:  "IX",
		5:  "X",
		6:  "XL",
		7:  "L",
		8:  "XC",
		9:  "C",
		10: "CD",
		11: "D",
		12: "CM",
		13: "M"}
}

func intToRoman(num int) string {
	var ans string;
	numMap := numMap()
	symbolsMap := symbolsMap()
	for i := 13; i > 0; i-- {
		if num >= numMap[i] {
			appendNum := num / numMap[i]
			num = num%numMap[i]
			for j := 0; j<appendNum; j++ {
				ans += symbolsMap[i]
			}
		}
	}
	return ans
}

func main()  {
	num := 1994
	fmt.Println(intToRoman(num))
}
