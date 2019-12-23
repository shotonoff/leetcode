package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumDecodings(t *testing.T) {
	assert.Equal(t, 2, numDecodings("12"))
	assert.Equal(t, 8, numDecodings("11111"))
	//assert.Equal(t, 2, numDecodings("115151"))
	//assert.Equal(t, 0, numDecodings(""))
	//assert.Equal(t, 0, numDecodings("0"))
}