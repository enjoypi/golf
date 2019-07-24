package queue

// go test -bench=. -gcflags "-N -l" ./...
import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var ()

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type QueueTestSuite struct {
	suite.Suite
}

// Make sure that Variable is set to five
// before each test
func (suite *QueueTestSuite) SetupTest() {
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
	size := 128
	q := NewCircleArrayQueue(size, false)
	require.True(q.Empty())
	for i := 1; i < size; i++ {
		require.False(q.Full())
		require.True(q.Push(i))
	}
	require.True(q.Full())
	require.False(q.Push(0))
	for i := 1; i < size; i++ {
		require.Equal(i, q.Pop())
	}
	require.Nil(q.Pop())
	require.True(q.Empty())
}

func doing(n int, q Queue) {
	for i := 0; i < n*2; i++ {
		if i%2 == 0 {
			q.Push(i)
		} else {
			q.Pop()
		}
	}
}

func BenchmarkCircleQueue(b *testing.B) {
	q := NewCircleArrayQueue(b.N, false)
	doing(b.N, q)
}

func BenchmarkSingleCircleQueue(b *testing.B) {
	q := NewCircleArrayQueue(b.N, true)
	doing(b.N, q)
}

func BenchmarkQueueWithLock(b *testing.B) {
	q := NewCircleArrayQueue(b.N, true)
	b.SetParallelism(4)
	b.RunParallel(func(pb *testing.PB) {
		doing(b.N, q)
		for pb.Next() {

		}
	})
}

func BenchmarkQueueChannel(b *testing.B) {
	q := NewChannelQueue(b.N)
	b.SetParallelism(4)
	b.RunParallel(func(pb *testing.PB) {
		doing(b.N, q)
		for pb.Next() {

		}
	})
}

func BenchmarkLockFreeQueue(b *testing.B) {
	q := NewLockFreeQueue(b.N)
	b.SetParallelism(4)
	b.RunParallel(func(pb *testing.PB) {
		doing(b.N, q)
		for pb.Next() {

		}
	})
}
