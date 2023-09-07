package leetcode

func longestConsecutive(nums []int) int {
	set := map[int]struct{}{}

	for _, n := range nums {
		set[n] = struct{}{}
	}

	longest := 0
	for _, n := range nums {
		len := seq_len(n, set, len(nums))
		if len > longest {
			longest = len
		}
	}

	return longest
}

func seq_len(n int, set map[int]struct{}, max int) int {
	_, prev_exists := set[n-1]
	if prev_exists {
		return 0
	}

	for i := 1; i <= max; i++ {
		_, next_exists := set[n+i]
		if !next_exists {
			return i
		}
	}
	panic("something's wrong with the set")
}
