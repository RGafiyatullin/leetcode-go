package leetcode

func generateParenthesis(n int) []string {
	out := []string{}
	gen(&out, "", n, 0)
	return out
}

func gen(out *[]string, prefix string, to_open int, to_close int) {
	if to_open > 0 {
		gen(out, prefix+"(", to_open-1, to_close+1)
	}
	if to_close > 0 {
		gen(out, prefix+")", to_open, to_close-1)
	}

	if to_open == 0 && to_close == 0 {
		*out = append(*out, prefix)
	}
}
