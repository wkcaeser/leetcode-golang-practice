package main

import "fmt"

func main() {
	//nums := []int {3, 2, 2, 3}
	nums := []int{3, 3}
	fmt.Println(removeElement(nums, 5))

	fmt.Println(nums)
}

/**
将目标值全部置换到数组末尾，统计末尾目标值的数量，数组长度直接减去目标值即可
*/

func removeElement(nums []int, val int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		if nums[0] == val {
			nums[0] = 0
			return 0
		} else {
			return 1
		}
	}

	endPos := len(nums) - 1

	for i, num := range nums {
		if i >= endPos {
			break
		}
		if num == val {
			for val == nums[endPos] && endPos > i {
				endPos--
			}
			nums[i], nums[endPos] = nums[endPos], nums[i]
		}
	}

	equalCnt := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			equalCnt++
		} else {
			break
		}
	}

	return len(nums) - equalCnt
}
