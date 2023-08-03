package main

func TwoSum1(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	seenNums := make(map[int]int)
	for index, thisNum := range nums {
		if seenIndex, ok := seenNums[target-thisNum]; ok {
			return []int{seenIndex, index}
		}
		seenNums[thisNum] = index
	}
	return []int{0, 0} // Should not happen
}

func twoSum3(nums []int, target int) []int {

	// dictionary
	// key: number
	// value: index
	num_idx_dict := make(map[int]int)

	// scan each (index, number) in input array nums
	for cur_idx, number := range nums {

		complement := target - number

		if idx_of_dual, dual_exist := num_idx_dict[complement]; dual_exist {

			// find two numbers with sum = target
			return []int{idx_of_dual, cur_idx}

		} else {

			// update current number with index into dictionary
			num_idx_dict[number] = cur_idx
		}
	}

	// Description guarantees that there is exactly one solution in online test, so this won't happen on run-time.
	return []int{-1, -1}
}
