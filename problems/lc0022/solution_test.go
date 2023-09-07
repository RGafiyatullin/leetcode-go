package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, []string{"()"}, generateParenthesis(1))
}
func TestExample2(t *testing.T) {
	assert.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, generateParenthesis(3))
}
