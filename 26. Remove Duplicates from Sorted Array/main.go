func removeDuplicates(nums []int) int {
	n := len(nums)

	if n <= 1 {
		return n
	}

	slow := 0
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}