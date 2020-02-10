package string

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://leetcode.com/problems/decode-ways/
// 91. Decode Ways
func numDecodings(s string) int {
	l := len(s)

	dp := make([]int, l+1)
	dp[0] = 1
	dp[1] = 1
	p := s[0] - 48

	for i := 2; i < l+1; i++ {
		c := s[i-1] - 48

		if (p*10)+c <= 26 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}

		p = c
	}

	return dp[l]
}

// https://leetcode.com/problems/roman-to-integer/
// 13. Roman to Integer
func romanToInt(s string) int {
	l := len(s)
	buf := make([]int, l)

	for i := 0; i < l; i++ {
		switch s[i] {
		case 'I':
			buf[i] = 1
		case 'V':
			buf[i] = 5
		case 'X':
			buf[i] = 10
		case 'L':
			buf[i] = 50
		case 'C':
			buf[i] = 100
		case 'D':
			buf[i] = 500
		case 'M':
			buf[i] = 1000
		}
	}

	res := 0
	var j int

	for i := 0; i < l; i++ {
		c := buf[i]

		for j = i + 1; j < l && buf[i] == buf[j]; j++ {
			c += buf[i]
		}

		if j >= l || buf[j] < buf[i] {
			res += c
			i = j - 1
		} else {
			res += buf[j] - c
			i = j
		}
	}

	return res
}

// https://leetcode.com/problems/longest-palindromic-substring/
// 5. Longest Palindromic Substring
func longestPalindrome(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	var (
		rs, r   string = "", s[0:1]
		rc, rsc int
	)

	for i := 0; i < l-1; i++ {
		rs, rsc = helperLongestPalindrome(s, i, i, l)

		if rsc > rc {
			r, rc = rs, rsc
		}

		rs, rsc = helperLongestPalindrome(s, i, i+1, l)

		if rsc > rc {
			r, rc = rs, rsc
		}
	}

	return r
}

func helperLongestPalindrome(s string, i, j, l int) (string, int) {
	for i > 0 && j < l-1 && s[i] == s[j] {
		i--
		j++
	}

	if s[i] != s[j] {
		i++
		j--
	}

	return s[i : j+1], j - i + 1
}

// https://leetcode.com/problems/zigzag-conversion/
// 6. ZigZag Conversion
//
// P   A   H   N
// A P L S I I G
// Y   I   R
//
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
func convert(s string, numRows int) string {
	l := len(s)
	d := numRows

	if d > 1 {
		d = numRows*2 - 2
	}

	var p, level int
	buf := bytes.NewBuffer([]byte{})

	for i := 0; i < l; i++ {
		if p >= l {
			level++
			p = level
		}

		buf.WriteByte(s[p])

		td := d - (level * 2)
		if (p+td) < l && level > 0 && level < (numRows-1) {
			buf.WriteByte(s[p+td])
			i++
		}

		p += d
	}

	return buf.String()
}

// https://leetcode.com/problems/palindromic-substrings/
// 647. Palindromic Substrings
func countSubstrings(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	var r = 0

	for i := 0; i < l; i++ {
		r += helperCountSubstrings(s, i, i, l)
		r += helperCountSubstrings(s, i, i+1, l)
	}

	return r
}

func helperCountSubstrings(s string, i, j, l int) (cnt int) {
	for i >= 0 && j < l && s[i] == s[j] {
		i--
		j++
		cnt++
	}

	return
}

// https://leetcode.com/problems/longest-palindromic-subsequence/
// 516. Longest Palindromic Subsequence
func longestPalindromeSubseq(s string) int {
	totalLen := len(s)
	if totalLen == 0 {
		return 0
	}

	dp := make([][]int, totalLen+1)
	for i := 0; i < totalLen+1; i++ {
		dp[i] = make([]int, totalLen)
		if i < totalLen {
			dp[i][i] = 1
		}
	}

	for partialLen := 1; partialLen < totalLen; partialLen++ {
		for i := 0; i < totalLen-partialLen; i++ {
			j := i + partialLen

			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	return dp[0][totalLen-1]
}

// https://leetcode.com/problems/longest-common-subsequence
// 1143. Longest Common Subsequence
func longestCommonSubsequence(text1, text2 string) int {
	l1 := len(text1)
	l2 := len(text2)

	dp := make([][]int, l1+1)

	for i := 0; i < l1+1; i++ {
		dp[i] = make([]int, l2+1)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = max(dp[i-1][j-1]+1, dp[i][j-1])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[l1][l2]
}

// https://leetcode.com/problems/length-of-last-word/
// 58. Length of Last Word
func lengthOfLastWord(s string) int {
	var i, j int
	l := len(s)

	for j = l - 1; j >= 0 && s[j] == ' '; j-- {
	}

	for i = j; i >= 0 && s[i] != ' '; i-- {
	}

	return j - i
}

func reverseVowels(s string) string {
	l := len(s)

	if l == 1 || l == 0 {
		return s
	}

	var i, j = 0, len(s) - 1
	r := make([]byte, len(s))

	for ; i <= j; i, j = i+1, j-1 {
		for ; i < j && !isVowel(s[i]); i++ {
			r[i] = s[i]
		}

		for ; j > i && !isVowel(s[j]); j-- {
			r[j] = s[j]
		}

		r[i], r[j] = s[j], s[i]
	}

	return string(r)
}

func isVowel(ch byte) bool {
	if ch == 'e' || ch == 'i' || ch == 'u' || ch == 'o' || ch == 'a' ||
		ch == 'E' || ch == 'I' || ch == 'U' || ch == 'O' || ch == 'A' {
		return true
	}

	return false
}

func reverseString(s []byte) {
	var i, j = 0, len(s) - 1

	for ; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// https://leetcode.com/problems/reverse-string-ii/submissions/
// 541. Reverse String II
func reverseStr(s string, k int) string {
	b := []byte(s)

	l := len(s)

	for i := 0; i < l; i = i + 2*k {
		j := i + k - 1

		if j >= l {
			j = l - 1
		}

		for l, r := i, j; l < r; l, r = l+1, r-1 {
			b[l], b[r] = b[r], b[l]
		}
	}

	return string(b)
}

// https://leetcode.com/problems/reverse-only-letters/
// 917. Reverse Only Letters
func reverseOnlyLetters(s string) string {
	b := []byte(s)
	var i, j = 0, len(s) - 1

	for ; i <= j; i, j = i+1, j-1 {
		for ; i < j && !isAlphabet(b[i]); i++ {
		}

		for ; i < j && !isAlphabet(b[j]); j-- {
		}

		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

func isAlphabet(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

// https://leetcode.com/problems/repeated-substring-pattern/
// 459. Repeated Substring Pattern
func repeatedSubstringPattern(s string) bool {
	l := len(s)

	if l <= 1 {
		return false
	}

	for i := 1; i <= l/2; i++ {
		if l%i == 0 && helperRepeatedSubstringPattern(i, l, s) {
			return true
		}
	}

	return false
}

func helperRepeatedSubstringPattern(j, l int, s string) bool {
	for i := 0; j < l; i, j = i+1, j+1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

// https://leetcode.com/problems/custom-sort-string/
// 791. Custom Sort String
func customSortString(s string, T string) string {
	ht := make([]int, 26)
	l := len(s)
	for i := 0; i < 26; i++ {
		ht[i] = 99
	}
	for i := 0; i < l; i++ {
		ht[s[i]-'a'] = i
	}

	res := []byte(T)
	sort.Slice(res, func(i, j int) bool {
		return ht[res[i]-'a'] < ht[res[j]-'a']
	})

	return string(res)
}

// https://leetcode.com/problems/simplify-path/
// 71. Simplify Path
func simplifyPath(path string) string {
	l := len(path)
	var j int
	var st []string
	var part string

	for i := 0; i < l; i = j + 1 {
		for j = i; j < l && path[j] != '/'; j++ {
		}

		part = path[i:j]
		if part == ".." {
			if len(st) > 0 {
				st = st[0 : len(st)-1]
			}
		} else if part != "." && part != "" {
			st = append(st, path[i:j])
		}
	}

	return fmt.Sprintf("/%s", strings.Join(st, "/"))
}
