package queue

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type QueueTestSuite struct {
	suite.Suite
	Queue
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *QueueTestSuite) SetupTest() {
	suite.Queue = NewQueue()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestQueueTestSuite(t *testing.T) {
	suite.Run(t, new(QueueTestSuite))
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *QueueTestSuite) TestNewQueue() {
	q := suite.Queue
	q.Push(1)
	q.Push(3)
	q.Push(9)
	require:= suite.Require()
	require.Equal(1, q.Pop())
	require.Equal(3, q.Pop())
	require.Equal(9, q.Pop())
	require.Equal(true, q.Empty())
}
