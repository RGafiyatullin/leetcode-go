package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, t := range tokens {
		val, err := strconv.Atoi(t)
		if err != nil {
			stack = pop_apply_push(stack, t)
		} else {
			stack = append(stack, val)
		}
	}

	if len(stack) != 1 {
		panic("bad stack state")
	}
	return stack[0]
}

func pop_apply_push(stack []int, op string) []int {
	right := stack[len(stack)-1]
	left := stack[len(stack)-2]
	result := 0

	switch op {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		result = left / right
	}

	out := stack[:len(stack)-2]
	return append(out, result)
}
