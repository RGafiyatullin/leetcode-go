package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, 4, search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func TestExample2(t *testing.T) {
	assert.Equal(t, -1, search([]int{-1, 0, 3, 5, 9, 12}, 2))
}

func TestCase1(t *testing.T) {
	assert.Equal(t, 0, search([]int{5}, 5))
}
