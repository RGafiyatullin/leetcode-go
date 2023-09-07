package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, 4, longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func TestExample2(t *testing.T) {
	assert.Equal(t, 9, longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
