package string

import (
	"bytes"
)

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
