package stack

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestStackTestSuite(t *testing.T) {
	suite.Run(t, new(StackTestSuite))
}

type StackTestSuite struct {
	suite.Suite
}

func (suite *StackTestSuite) SetupTest() {

}

func (suite *StackTestSuite) TestMinStack() {
	st := Constructor()

	st.Push(6)
	assert.Equal(suite.T(), 6, st.GetMin())
	assert.Equal(suite.T(), 6, st.Top())

	st.Push(4)
	assert.Equal(suite.T(), 4, st.GetMin())
	assert.Equal(suite.T(), 4, st.Top())

	st.Push(10)
	assert.Equal(suite.T(), 4, st.GetMin())
	assert.Equal(suite.T(), 10, st.Top())

	st.Push(1)
	assert.Equal(suite.T(), 1, st.GetMin())
	assert.Equal(suite.T(), 1, st.Top())

	st.Push(2)
	assert.Equal(suite.T(), 1, st.GetMin())
	assert.Equal(suite.T(), 2, st.Top())

	st.Push(3)
	assert.Equal(suite.T(), 1, st.GetMin())
	assert.Equal(suite.T(), 3, st.Top())

	st.Pop()
	assert.Equal(suite.T(), 1, st.GetMin())
	assert.Equal(suite.T(), 2, st.Top())

	st.Pop()
	st.Pop()
	assert.Equal(suite.T(), 4, st.GetMin())
	assert.Equal(suite.T(), 10, st.Top())

	st.Pop()
	st.Pop()
	assert.Equal(suite.T(), 6, st.GetMin())
	assert.Equal(suite.T(), 6, st.Top())
}
