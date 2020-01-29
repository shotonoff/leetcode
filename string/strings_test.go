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

func (suite *StringTestSuite) TestTwoSum() {
	assert.Equal(suite.T(), []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(suite.T(), []int(nil), twoSum([]int{}, 9))
	assert.Equal(suite.T(), []int{0, 3}, twoSum([]int{2, 23, 11, 7}, 9))
	assert.Equal(suite.T(), []int{1, 2}, twoSum([]int{3, 2, 4}, 6))
	assert.Equal(suite.T(), []int{0, 1}, twoSum([]int{3, 3}, 6))
}

func (suite *StringTestSuite) TestTwoSum2() {
	assert.Equal(suite.T(), []int{1, 2}, twoSum2([]int{2, 7, 11, 15}, 9))
	assert.Equal(suite.T(), []int{2, 4}, twoSum2([]int{2, 7, 11, 23}, 30))
	assert.Equal(suite.T(), []int{4, 7}, twoSum2([]int{2, 3, 4, 5, 6, 7, 8, 100}, 13))
	assert.Equal(suite.T(), []int{1, 2}, twoSum2([]int{3, 3}, 6))
	assert.Equal(suite.T(), []int{3, 4}, twoSum2([]int{1, 2, 3, 3, 7, 8, 9}, 6))
}

func TestS3ClientTestSuite(t *testing.T) {
	suite.Run(t, new(StringTestSuite))
}
