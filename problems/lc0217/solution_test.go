package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, containsDuplicate([]int{1, 2, 3, 1}), true)
}

func TestExample2(t *testing.T) {
	assert.Equal(t, containsDuplicate([]int{1, 2, 3, 4}), false)
}

func TestExample3(t *testing.T) {
	assert.Equal(t, containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}), true)
}
