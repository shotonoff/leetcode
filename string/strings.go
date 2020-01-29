package string

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

// https://leetcode.com/problems/two-sum/
// 1. Two Sum
func twoSum(nums []int, target int) []int {
	ht := make(map[int]int)

	for i, n := range nums {
		m := target - n
		if j, ok := ht[m]; ok == true {
			return []int{j, i}
		}

		ht[n] = i
	}

	return nil
}

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// 167. Two Sum II - Input array is sorted
func twoSum2(nums []int, target int) []int {
	var i, j, sum = 0, len(nums) - 1, 0

	for i < j {
		sum = nums[i] + nums[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}

		if sum > target {
			j--
		} else {
			i++
		}
	}

	return nil
}
