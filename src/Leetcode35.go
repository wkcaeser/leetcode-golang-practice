package main

func searchInsert(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	begin := 0
	end := l

	for begin < end {
		mid := (begin + end) / 2
		if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			begin = mid + 1
		} else {
			return mid
		}
	}

	return begin
}
