package leetcode

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
