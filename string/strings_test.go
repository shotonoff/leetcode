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

func TestS3ClientTestSuite(t *testing.T) {
	suite.Run(t, new(StringTestSuite))
}
