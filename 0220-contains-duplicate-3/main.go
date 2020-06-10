package main

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) (res bool) {
	lenNums := len(nums)
	for  i := 0; i < lenNums; i++ {
		for j := i + 1; j < lenNums; j++ {
			if abs(nums[i] - nums[j]) <= t && abs(i - j) <= k {
				return true
			}
		}
	}
	return false
}
