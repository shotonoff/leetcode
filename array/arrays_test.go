package array

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestArrayTestSuite(t *testing.T) {
	suite.Run(t, new(ArrayTestSuite))
}

type ArrayTestSuite struct {
	suite.Suite
}

func (suite *ArrayTestSuite) TestTwoSum() {
	assert.Equal(suite.T(), []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(suite.T(), []int(nil), twoSum([]int{}, 9))
	assert.Equal(suite.T(), []int{0, 3}, twoSum([]int{2, 23, 11, 7}, 9))
	assert.Equal(suite.T(), []int{1, 2}, twoSum([]int{3, 2, 4}, 6))
	assert.Equal(suite.T(), []int{0, 1}, twoSum([]int{3, 3}, 6))
}

func (suite *ArrayTestSuite) TestTwoSum2() {
	assert.Equal(suite.T(), []int{1, 2}, twoSum2([]int{2, 7, 11, 15}, 9))
	assert.Equal(suite.T(), []int{2, 4}, twoSum2([]int{2, 7, 11, 23}, 30))
	assert.Equal(suite.T(), []int{4, 7}, twoSum2([]int{2, 3, 4, 5, 6, 7, 8, 100}, 13))
	assert.Equal(suite.T(), []int{1, 2}, twoSum2([]int{3, 3}, 6))
	assert.Equal(suite.T(), []int{3, 4}, twoSum2([]int{1, 2, 3, 3, 7, 8, 9}, 6))
}

func (suite *ArrayTestSuite) TestThreeSum() {
	var expected [][]int

	expected = [][]int{{-1, -1, 2}, {-1, 0, 1}}
	assert.Equal(suite.T(), expected, threeSum([]int{-1, 0, 1, 2, -1, -4}))

	expected = [][]int{{0, 0, 0}}
	assert.Equal(suite.T(), expected, threeSum([]int{0, 0, 0, 0, 0, 0, 0, 0}))

	expected = [][]int{{-2, 0, 2}, {-2, 1, 1}, {-1, 0, 1}, {0, 0, 0}}
	assert.Equal(suite.T(), expected, threeSum([]int{0, 1, -1, 0, 1, -2, 1, 2, 1, 0, 0}))

	expected = [][]int{{-2, 1, 1}}
	assert.Equal(suite.T(), expected, threeSum([]int{1, 1, -2}))

	expected = [][]int{{0, 0, 0}}
	assert.Equal(suite.T(), expected, threeSum([]int{0, 0, 0}))

	expected = [][]int{{-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}}
	assert.Equal(suite.T(), expected, threeSum([]int{3, 0, -2, -1, 1, 2}))

	expected = [][]int{{-2, 0, 2}}
	assert.Equal(suite.T(), expected, threeSum([]int{-2, 0, 0, 2, 2}))

	expected = [][]int{{-1, 0, 1}}
	assert.Equal(suite.T(), expected, threeSum([]int{1, -1, -1, 0}))

	expected = [][]int(nil)
	assert.Equal(suite.T(), expected, threeSum([]int{-1, -2, -1, -1, -1, -2, -3, -4}))
}

func (suite *ArrayTestSuite) TestThreeSumClosest() {
	assert.Equal(suite.T(), 2, threeSumClosest([]int{-1, 2, 1, -4}, 1))
	assert.Equal(suite.T(), 1, threeSumClosest([]int{1, 1, -1, -1, 3}, 1))
	assert.Equal(suite.T(), 2, threeSumClosest([]int{-1, 0, 1, 1, 55}, 3))

	r := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	for i, f := range r {
		assert.Equal(suite.T(), i, helperThreeSumClosest(r, f, 0, len(r)-1))
	}

	r = []int{-1, 0, 1, 1, 55}
	assert.Equal(suite.T(), 3, helperThreeSumClosest(r, 2, 3, len(r)-1))
}

func (suite *ArrayTestSuite) TestMerge() {
	var actual, expected [][]int

	actual = [][]int{}
	assert.Equal(suite.T(), [][]int{}, merge(actual))

	actual = [][]int{{1, 4}, {1, 4}}
	expected = [][]int{{1, 4}}
	assert.Equal(suite.T(), expected, merge(actual))

	actual = [][]int{{1, 4}, {4, 5}}
	expected = [][]int{{1, 5}}
	assert.Equal(suite.T(), expected, merge(actual))

	actual = [][]int{{1, 3}, {5, 7}}
	expected = [][]int{{1, 3}, {5, 7}}
	assert.Equal(suite.T(), expected, merge(actual))

	actual = [][]int{{1, 3}}
	expected = [][]int{{1, 3}}
	assert.Equal(suite.T(), expected, merge(actual))

	actual = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	expected = [][]int{{1, 6}, {8, 10}, {15, 18}}
	assert.Equal(suite.T(), expected, merge(actual))

	actual = [][]int{{1, 3}, {2, 6}, {8, 10}, {13, 22}, {15, 18}}
	expected = [][]int{{1, 6}, {8, 10}, {13, 22}}
	assert.Equal(suite.T(), expected, merge(actual))

	actual = [][]int{{1, 11}, {2, 6}, {8, 10}, {13, 22}, {15, 18}}
	expected = [][]int{{1, 11}, {13, 22}}
	assert.Equal(suite.T(), expected, merge(actual))

	actual = [][]int{{1, 2}, {3, 6}, {8, 10}, {13, 14}, {15, 18}}
	expected = [][]int{{1, 2}, {3, 6}, {8, 10}, {13, 14}, {15, 18}}
	assert.Equal(suite.T(), expected, merge(actual))

	actual = [][]int{{1, 20}, {2, 6}, {8, 10}, {13, 22}, {15, 18}}
	expected = [][]int{{1, 22}}
	assert.Equal(suite.T(), expected, merge(actual))
}

func (suite *ArrayTestSuite) TestMaxSlidingWindow() {
	assert.Equal(suite.T(), []int{3, 3, 5, 5, 6, 7}, maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	assert.Equal(suite.T(), []int{3, 3, 10, 10, 10, 6, 7}, maxSlidingWindow([]int{1, 3, -1, -3, 10, 3, 6, 1, 7}, 3))
	assert.Equal(suite.T(), []int{}, maxSlidingWindow([]int{}, 0))
	assert.Equal(suite.T(), []int{1}, maxSlidingWindow([]int{1}, 0))

	a := []int{-95, 92, -85, 59, -59, -14, 88, -39, 2, 92, 94, 79, 78, -58, 37, 48, 63, -91, 91, 74, -28, 39, 90, -9, -72, -88, -72, 93, 38, 14, -83, -2, 21, 4, -75, -65, 3, 63, 100, 59, -48, 43, 35, -49, 48, -36, -64, -13, -7, -29, 87, 34, 56, -39, -5, -27, -28, 10, -57, 100, -43, -98, 19, -59, 78, -28, -91, 67, 41, -64, 76, 5, -58, -89, 83, 26, -7, -82, -32, -76, 86, 52, -6, 84, 20, 51, -86, 26, 46, 35, -23, 30, -51, 54, 19, 30, 27, 80, 45, 22}
	e := []int{92, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 91, 91, 91, 91, 91, 91, 91, 93, 93, 93, 93, 93, 93, 93, 93, 93, 93, 63, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 59, 48, 87, 87, 87, 87, 87, 87, 87, 87, 87, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 78, 78, 78, 78, 78, 83, 83, 83, 83, 83, 83, 86, 86, 86, 86, 86, 86, 86, 86, 86, 86, 84, 84, 84, 54, 54, 54, 54, 80, 80, 80}
	assert.Equal(suite.T(), e, maxSlidingWindow(a, 10))
}
