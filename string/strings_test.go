package string

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type StringTestSuite struct {
	suite.Suite
}

func (suite *StringTestSuite) SetupTest() {

}

func (suite *StringTestSuite) TestNumDecodings() {
	assert.Equal(suite.T(), 2, numDecodings("12"))
	assert.Equal(suite.T(), 8, numDecodings("11111"))
	//assert.Equal(suite.T(), 2, numDecodings("115151"))
	//assert.Equal(suite.T(), 0, numDecodings(""))
	//assert.Equal(suite.T(), 0, numDecodings("0"))
}

func (suite *StringTestSuite) TestRomanToInt() {
	assert.Equal(suite.T(), 0, romanToInt(""))
	assert.Equal(suite.T(), 3, romanToInt("III"))
	assert.Equal(suite.T(), 4, romanToInt("IV"))
	assert.Equal(suite.T(), 9, romanToInt("IX"))
	assert.Equal(suite.T(), 58, romanToInt("LVIII"))
	assert.Equal(suite.T(), 1994, romanToInt("MCMXCIV"))
	assert.Equal(suite.T(), 1994, romanToInt("XXXIX"))
	assert.Equal(suite.T(), 1994, romanToInt("CCXLVI"))
}

func (suite *StringTestSuite) TestConvert() {
	assert.Equal(suite.T(), "PAYPALISHIRING", convert("PAYPALISHIRING", 1))
	assert.Equal(suite.T(), "PYAIHRNAPLSIIG", convert("PAYPALISHIRING", 2))
	assert.Equal(suite.T(), "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
	assert.Equal(suite.T(), "PINALSIGYAHRPI", convert("PAYPALISHIRING", 4))
	assert.Equal(suite.T(), "PHASIYIRPLIGAN", convert("PAYPALISHIRING", 5))
}

func (suite *StringTestSuite) TestLongestPalindrome() {
	assert.Equal(suite.T(), "a", longestPalindrome("a"))
	assert.Equal(suite.T(), "a", longestPalindrome("ab"))
	assert.Equal(suite.T(), "aa", longestPalindrome("aa"))
	assert.Equal(suite.T(), "aa", longestPalindrome("aab"))
	assert.Equal(suite.T(), "aa", longestPalindrome("baa"))
	assert.Equal(suite.T(), "bab", longestPalindrome("babad"))
	assert.Equal(suite.T(), "bb", longestPalindrome("cbbd"))
	assert.Equal(suite.T(), "asdfdsa", longestPalindrome("asdfdsa"))
	assert.Equal(suite.T(), "asdfdsa", longestPalindrome("dasdfdsal"))
	assert.Equal(suite.T(), "ff", longestPalindrome("asdff"))
	assert.Equal(suite.T(), "ddddddd", longestPalindrome("asdffddddddd"))
	assert.Equal(suite.T(), "dddddddd", longestPalindrome("ddddddddasdffddddddd"))
	assert.Equal(suite.T(), "a", longestPalindrome("asdfghjkl"))
}

func TestS3ClientTestSuite(t *testing.T) {
	suite.Run(t, new(StringTestSuite))
}
