package leetcode

func search(nums []int, target int) int {

	lo := 0
	hi := len(nums) - 1

	for lo <= hi {
		mid := (lo + hi) / 2
		val := nums[mid]

		// fmt.Println("lo=", lo, " hi=", hi, " mid=", mid, " value=", nums[mid])

		if val == target {
			return mid
		} else if val < target {
			lo = mid + 1
		} else if val > target {
			hi = mid - 1
		}
	}

	return -1
}
