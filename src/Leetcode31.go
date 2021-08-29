package main

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			minPos := i
			for j := i; j < len(nums); j++ {
				if nums[j] > nums[i-1] {
					minPos = j
				}
			}
			nums[i-1], nums[minPos] = nums[minPos], nums[i-1]

			reverse(nums[i:])
			return
		}
	}

	reverse(nums)
}

func reverse(nums []int) {
	l := len(nums)

	for i := 0; i < l/2; i++ {
		nums[i], nums[l-i-1] = nums[l-i-1], nums[i]
	}
}
