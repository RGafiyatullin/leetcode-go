package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1})
}

func TestExample2(t *testing.T) {
	assert.Equal(t, twoSum([]int{3, 2, 4}, 6), []int{1, 2})
}

func TestExample3(t *testing.T) {
	assert.Equal(t, twoSum([]int{3, 3}, 6), []int{0, 1})
}

func TestCase1(t *testing.T) {
	assert.Equal(t, twoSum([]int{3, 2, 4}, 6), []int{1, 2})
}
