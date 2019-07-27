package queue

// go test -bench=. -benchtime 200ms -gcflags "-N -l" ./...
import (
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/suite"
)

var (
	gQ  = NewCircleArrayQueue(20000000, false)
	gQL = NewCircleArrayQueue(20000000, true)
	gSQ = NewSliceQueue(2000000)
)

func init() {
	for i := 0; i < 20000000; i++ {
		gQ.Push(i)
		gQL.Push(i)
		gSQ.Push(i)
	}
	logrus.Info(gQ.Size())
}

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
func (suite *QueueTestSuite) TestSliceQueue() {
	size := 1024
	q := NewSliceQueue(size)
	testQueue(&suite.Suite, q, size)
}

func (suite *QueueTestSuite) TestCircleQueue() {
	size := 1024
	q := NewCircleArrayQueue(size, false)
	testQueue(&suite.Suite, q, size)
}
func (suite *QueueTestSuite) TestCircleQueueWithLock() {
	size := 512
	q := NewCircleArrayQueue(size, true)
	testQueue(&suite.Suite, q, size)
}
func (suite *QueueTestSuite) TestChannelQueue() {
	size := 256
	q := NewChannelQueue(size)
	testQueue(&suite.Suite, q, size)
}
func (suite *QueueTestSuite) TestLockFreeQueue() {
	size := 128
	q := NewLockFreeQueue(size)
	testQueue(&suite.Suite, q, size)
}

func testQueue(suite *suite.Suite, q Queue, size int) {
	require := suite.Require()
	require.True(q.Empty())
	for i := 0; i < size; i++ {
		require.False(q.Full())
		require.True(q.Push(i))
	}
	suite.T().Log(q.Size())
	require.True(q.Full())
	require.False(q.Push(0))
	for i := 0; i < size; i++ {
		require.Equal(i, q.Pop())
	}
	require.Nil(q.Pop())
	require.True(q.Empty())
}

// Benchmark
func doing(n int, q Queue) {
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			q.Push(i)
		} else {
			q.Pop()
		}
	}
}

func BenchmarkSliceQueuePop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gSQ.Pop()
	}
}

func BenchmarkSliceQueuePush(b *testing.B) {
	q := NewSliceQueue(b.N)
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}

func BenchmarkCircleQueuePop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gQ.Pop()
	}
}

func BenchmarkCircleQueuePush(b *testing.B) {
	q := NewCircleArrayQueue(b.N, false)
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}

func BenchmarkCircleQueueWithLockPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gQL.Pop()
	}
}

func BenchmarkCircleQueueWithLockPush(b *testing.B) {
	q := NewCircleArrayQueue(b.N, true)
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}

func doNext(i int, q Queue) int {
	if i%2 == 0 {
		q.Push(i)
	} else {
		q.Pop()
	}
	return i + 1
}

func BenchmarkSliceQueue(b *testing.B) {
	q := NewSliceQueue(b.N)
	b.SetParallelism(4)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			i = doNext(i, q)
		}
	})
}

func BenchmarkCircleQueueWithLock(b *testing.B) {
	q := NewCircleArrayQueue(b.N, true)
	b.SetParallelism(4)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			i = doNext(i, q)
		}
	})
}

func _BenchmarkQueueChannel(b *testing.B) {
	q := NewChannelQueue(b.N)
	b.SetParallelism(4)
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			i = doNext(i, q)
		}
	})
}

func _BenchmarkLockFreeQueue(b *testing.B) {
	q := NewLockFreeQueue(b.N)
	b.SetParallelism(4)
	b.RunParallel(func(pb *testing.PB) {
		doing(b.N, q)
		for pb.Next() {

		}
	})
}
