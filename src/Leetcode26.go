package main

/**
方法1：
 数组有序
 1、如果数据长度小于2 则只有1或者0个元素，返回数组长度即可，数组本身不需要操作
 2、从第二个元素开始遍历数组，如果后一个数据与前一个数据不同，则元素个数加一
 3、当后一个元素与前一个元素不相同时，将前面的元素相同的部分置为后面元素的值，则最后，在前面的几个元素则为不同元素
	例    1 1 2 2 2 3 4 4 5 5
		1 |1 2 2 2 3 4 4 5 5    1
		1 2 |2 2 2 3 4 4 5 5    2
		1 2 2 |2 2 3 4 4 5 5    2
		1 2 2 2 |2 3 4 4 5 5    2
		1 2 3 3 3 |3 4 4 5 5    3
		1 2 3 4 4 4 |4 4 5 5    4
		1 2 3 4 4 4 4 |4 5 5    4
		1 2 3 4 5 5 5 5 |5 5    5
		1 2 3 4 5 5 5 5 5 |5    5
思路2：
  在1的基础上记录当前已有的不同数的个数，直接置换
*/
//func removeDuplicates(nums []int) int {
//	if len(nums) < 2 {
//		return len(nums)
//	}
//
//	rs := 1
//	for i:=1; i < len(nums); i++ {
//		if nums[i] != nums[i-1] {
//			rs ++
//			for j := i-1; j > 0; j--  {
//				if nums[j] == nums[j-1] {
//					nums[j] = nums[i]
//				}
//			}
//		}
//	}
//
//	return rs
//}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	rs := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[rs] = nums[i]
			rs++
		}
	}

	return rs
}
