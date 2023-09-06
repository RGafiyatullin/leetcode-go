package leetcode

func containsDuplicate(nums []int) bool {
	hits := map[int]bool{}
	for _, num := range nums {
		if hits[num] {
			return true
		}
		hits[num] = true
	}
	return false
}
