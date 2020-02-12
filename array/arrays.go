package array

import (
	"math"
	"sort"
)

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

// https://leetcode.com/problems/3sum/
// 15. 3Sum
func threeSum(nums []int) (result [][]int) {
	if len(nums) < 3 {
		return
	}

	sort.Ints(nums)

	l := len(nums)
	var i, j int
	k := 0
	m := 0

	for ; m < l-2 && nums[m] <= 0; m++ {
		i = m + 1
		j = l - 1

		for i < j {
			sum3 := nums[m] + nums[i] + nums[j]

			if sum3 == k {
				y := []int{nums[m], nums[i], nums[j]}

				result = append(result, y)
				ki, kj := i, j

				for ; i < j && nums[ki] == nums[i] && nums[kj] == nums[j]; i, j = i+1, j-1 {
				}

			} else if sum3 > k {
				j--
			} else {
				i++
			}
		}

		for ; m < l-1 && nums[m] == nums[m+1]; m++ {
		}
	}

	return
}

// https://leetcode.com/problems/3sum-closest/
// 16. 3Sum Closest
func threeSumClosest(nums []int, t int) int {
	l := len(nums)

	if l < 3 {
		return 0
	}

	c := math.MaxInt32
	sum3 := math.MaxInt32

	sort.Ints(nums)

	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			r := helperThreeSumClosest(nums, t-(nums[i]+nums[j]), j+1, l-1)
			sum := nums[i] + nums[j] + nums[r]
			g := int(math.Abs(float64(t - sum)))

			if g < c {
				c = g
				sum3 = sum
			}
		}
	}

	return sum3
}

func helperThreeSumClosest(nums []int, t, l, r int) int {
	m := l + (r-l)/2

	if m-l > 0 && t < nums[m] {
		r := helperThreeSumClosest(nums, t, l, m-1)

		if int(math.Abs(float64(t-nums[r]))) < int(math.Abs(float64(t-nums[m]))) {
			return r
		}
	}

	if r-m > 0 && t > nums[m] {
		r := helperThreeSumClosest(nums, t, m+1, r)

		if int(math.Abs(float64(t-nums[r]))) < int(math.Abs(float64(t-nums[m]))) {
			return r
		}
	}

	return m
}

// https://leetcode.com/problems/merge-intervals/
// 56. Merge Intervals
func merge(in [][]int) (res [][]int) {
	l := len(in)
	if l <= 1 {
		return in
	}

	sort.Slice(in, func(i, j int) bool {
		return in[i][0] < in[j][0]
	})

	var n = in[0]

	for i := 1; i < l; i++ {
		if n[1] < in[i][0] {
			res = append(res, n)
			n = in[i]
		} else if n[1] >= in[i][0] {
			n[1] = Max(n[1], in[i][1])
		}
	}

	return append(res, n)
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
