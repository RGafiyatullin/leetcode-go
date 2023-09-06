package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, true, isValid("()"))
}

func TestExample2(t *testing.T) {
	assert.Equal(t, true, isValid("()[]{}"))
}

func TestExample3(t *testing.T) {
	assert.Equal(t, false, isValid("(]"))
}
