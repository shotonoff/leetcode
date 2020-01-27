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

func TestRomanToInt(t *testing.T) {
	assert.Equal(t, 0, romanToInt(""))
	assert.Equal(t, 3, romanToInt("III"))
	assert.Equal(t, 4, romanToInt("IV"))
	assert.Equal(t, 9, romanToInt("IX"))
	assert.Equal(t, 58, romanToInt("LVIII"))
	assert.Equal(t, 1994, romanToInt("MCMXCIV"))
	assert.Equal(t, 1994, romanToInt("XXXIX"))
	assert.Equal(t, 1994, romanToInt("CCXLVI"))
}
