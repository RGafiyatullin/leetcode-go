package leetcode

func maxProfit(prices []int) int {
	max := 0
	ds := make([]int, len(prices))

	for i := len(prices) - 1; i >= 0; i-- {
		ds[i] = max - prices[i]

		max = max_of_two(max, prices[i])
	}

	ret := max_of_two(0, find_max(ds))

	return ret
}

func find_max(ds []int) int {
	max := 0
	is_set := false

	for _, d := range ds {
		if !is_set {
			max = d
			is_set = true
		} else {
			max = max_of_two(max, d)
		}
	}

	return max
}

func max_of_two(a, b int) int {
	if a > b {
		return a
	}
	return b
}
