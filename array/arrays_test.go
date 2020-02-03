package array

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

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

func TestArrayTestSuite(t *testing.T) {
	suite.Run(t, new(ArrayTestSuite))
}
