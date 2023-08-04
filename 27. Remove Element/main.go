func removeElement(nums []int, val int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	slow := 0
	for fast := 0; fast < n; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

