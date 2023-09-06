package leetcode

func twoSum(nums []int, target int) []int {
	value_to_index := map[int]int{}

	for i, n := range nums {
		complement := target - n
		complement_idx, exists := value_to_index[complement]
		if exists {
			return []int{complement_idx, i}
		}

		value_to_index[n] = i
	}

	panic("at least one answer is guaranteed")
}
