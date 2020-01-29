package binaryTree

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// BTSSuite
type BTSSuite struct {
	suite.Suite
}

func (suite *BTSSuite) SetupTest() {

}

func (suite BTSSuite) TestFindTarget() {
	var root *TreeNode

	root = binaryTreeFromSlice([]interface{}{5, 3, 6, 2, 4, nil, 7})
	assert.Equal(suite.T(), true, findTarget(root, 9))
	assert.Equal(suite.T(), true, findTarget(root, 8))
	assert.Equal(suite.T(), true, findTarget(root, 11))
	assert.Equal(suite.T(), false, findTarget(root, 28))

	root = binaryTreeFromSlice([]interface{}{})
	assert.Equal(suite.T(), false, findTarget(root, 0))
}

func TestBTSSuite(t *testing.T) {
	suite.Run(t, new(BTSSuite))
}
