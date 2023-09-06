package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, maxProfit([]int{7, 1, 5, 3, 6, 4}), 5)
}

func TestExample2(t *testing.T) {
	assert.Equal(t, maxProfit([]int{7, 6, 4, 3, 1}), 0)
}
