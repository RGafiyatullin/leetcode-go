package leetcode

func isValid(s string) bool {
	stack := []rune{}

	for _, rune := range s {
		switch rune {
		case ')':
			if !pop(&stack, '(') {
				return false
			}
		case '}':
			if !pop(&stack, '{') {
				return false
			}
		case ']':
			if !pop(&stack, '[') {
				return false
			}

		default:
			stack = append(stack, rune)
		}
	}

	return len(stack) == 0
}

func pop(stack *[]rune, expected rune) bool {
	if len(*stack) == 0 {
		return false
	}
	last := (*stack)[len(*stack)-1]

	if last != expected {
		return false
	}
	*stack = (*stack)[:len(*stack)-1]
	return true
}
