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

func (suite *StringTestSuite) TestCountSubstrings() {
	assert.Equal(suite.T(), 1, countSubstrings("a"))
	assert.Equal(suite.T(), 2, countSubstrings("ab"))
	assert.Equal(suite.T(), 3, countSubstrings("aa"))
	assert.Equal(suite.T(), 3, countSubstrings("abc"))
	assert.Equal(suite.T(), 5, countSubstrings("abcb"))
	assert.Equal(suite.T(), 6, countSubstrings("aaa"))
}

func (suite *StringTestSuite) TestLongestPalindromeSubseq() {
	assert.Equal(suite.T(), 4, longestPalindromeSubseq("bbbab"))
	assert.Equal(suite.T(), 4, longestPalindromeSubseq("bbbsdfb"))
	assert.Equal(suite.T(), 5, longestPalindromeSubseq("bbbaaab"))
	assert.Equal(suite.T(), 2, longestPalindromeSubseq("cbbd"))
	assert.Equal(suite.T(), 2, longestPalindromeSubseq("aab"))
	assert.Equal(suite.T(), 1, longestPalindromeSubseq("a"))
}

func (suite *StringTestSuite) TestLongestCommonSubsequence() {
	assert.Equal(suite.T(), 6, longestCommonSubsequence("asddfghjkl", "dqwahyrdopfqqgrtkwel"))
	assert.Equal(suite.T(), 2, longestCommonSubsequence("addghl", "dahrdo"))

	assert.Equal(suite.T(), 2, longestCommonSubsequence("asd", "adqaydad"))
	assert.Equal(suite.T(), 3, longestCommonSubsequence("abcde", "ace"))
	assert.Equal(suite.T(), 3, longestCommonSubsequence("abc", "abc"))
	assert.Equal(suite.T(), 0, longestCommonSubsequence("abc", "ghj"))
	assert.Equal(suite.T(), 4, longestCommonSubsequence("pmjghexybyrgzczy", "hafcdqbgncrcbihkd"))
	assert.Equal(suite.T(), 3, longestCommonSubsequence("abcaf", "bkamcnf"))
	assert.Equal(suite.T(), 1, longestCommonSubsequence("abcaf", "a"))
	assert.Equal(suite.T(), 1, longestCommonSubsequence("a", "abcaf"))
}

func (suite *StringTestSuite) TestLengthOfLastWord() {
	assert.Equal(suite.T(), 4, lengthOfLastWord("asd asd qazs"))
	assert.Equal(suite.T(), 4, lengthOfLastWord("asd asd qazs   "))
	assert.Equal(suite.T(), 4, lengthOfLastWord("   asd asd qazs  "))
	assert.Equal(suite.T(), 2, lengthOfLastWord("qa zs  "))
	assert.Equal(suite.T(), 2, lengthOfLastWord("zs  "))
	assert.Equal(suite.T(), 2, lengthOfLastWord(" a s d f g g s_"))
	assert.Equal(suite.T(), 0, lengthOfLastWord("  "))
	assert.Equal(suite.T(), 5, lengthOfLastWord("Hello World"))
}

func (suite *StringTestSuite) TestReverseVowels() {
	assert.Equal(suite.T(), "holle", reverseVowels("hello"))
	assert.Equal(suite.T(), "leotcede", reverseVowels("leetcode"))
	assert.Equal(suite.T(), "", reverseVowels(""))
	assert.Equal(suite.T(), "a", reverseVowels("a"))
	assert.Equal(suite.T(), "ua", reverseVowels("au"))
	assert.Equal(suite.T(), "aA", reverseVowels("Aa"))
	assert.Equal(suite.T(), "aA", reverseVowels("Aa"))
	assert.Equal(suite.T(), "a a", reverseVowels("a a"))
	assert.Equal(suite.T(), "ab ba", reverseVowels("ab ba"))
	assert.Equal(suite.T(), "abA aba", reverseVowels("aba Aba"))
	assert.Equal(suite.T(), "abA  aba", reverseVowels("aba  Aba"))
}

func TestS3ClientTestSuite(t *testing.T) {
	suite.Run(t, new(StringTestSuite))
}
