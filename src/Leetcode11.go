package main

import (
	"fmt"
)

func minInt(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxArea(height []int) int {
	length := len(height)
	start, end := 0, length-1
	max := 0
	for ; start < end; {
		max = maxInt(max, (end - start) * minInt(height[end], height[start]))
		if height[end] < height[start] {
			end--
		} else {
			start++
		}
	}
	return max
}

func main() {
	heightPtr := &[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(*heightPtr))
}
