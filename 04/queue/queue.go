package queue

// Queue 队列
type Queue[T any] struct {
	data  []T
	front int
}

const DefaultCap = 100

// EmptyInfo 为空时的错误提示信息
const EmptyInfo = "queue is empty"

func NewQueue[T any]() *Queue[T] {
	return NewQueueWithCap[T](DefaultCap)
}

func NewQueueWithCap[T any](cap int) *Queue[T] {

	return &Queue[T]{make([]T, 0, cap), 0}
}

// IsEmpty 判断队列是否为空
func (q *Queue[T]) IsEmpty() bool {

	return q.front == len(q.data)
}

// EnQueue 入队
func (q *Queue[T]) EnQueue(v T) {

	q.data = append(q.data, v)
}

// DeQueue 出队
func (q *Queue[T]) DeQueue() T {

	if q.IsEmpty() {
		panic(EmptyInfo)
	}

	ret := q.data[q.front]

	q.front++

	return ret
}

// Front 获取队头元素
func (q *Queue[T]) Front() T {

	if q.IsEmpty() {
		panic(EmptyInfo)
	}

	return q.data[q.front]
}
