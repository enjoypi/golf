package queue

// go test -bench=. -gcflags "-N -l" ./...
import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	gQueue = NewQueueWithLock()
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type QueueTestSuite struct {
	suite.Suite
	Queue
}

// Make sure that Variable is set to five
// before each test
func (suite *QueueTestSuite) SetupTest() {
	suite.Queue = NewCircleArrayQueue(4096)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestQueueTestSuite(t *testing.T) {
	suite.Run(t, new(QueueTestSuite))
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *QueueTestSuite) TestNewQueue() {
	require := suite.Require()
	q := suite.Queue
	require.True(q.Empty())
	q.Push(1)
	q.Push(3)
	q.Push(9)
	require.Equal(1, q.Pop())
	require.Equal(3, q.Pop())
	require.Equal(9, q.Pop())
	require.Equal(true, q.Empty())
}

// Benchmark
func BenchmarkQueueTestSuite(b *testing.B) {
	bench := func(pb *testing.PB) {
		for i := 0; i < b.N; i++ {
			gQueue.Push(i)
		}
		for i := 0; i < b.N; i++ {
			gQueue.Pop()
		}
		for pb.Next() {

		}
	}
	b.SetParallelism(8)
	b.RunParallel(bench)
}
