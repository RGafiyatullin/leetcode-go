package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	assert.Equal(t, isAnagram("anagram", "nagaram"), true)
}

func TestExample2(t *testing.T) {
	assert.Equal(t, isAnagram("car", "rat"), false)
}
