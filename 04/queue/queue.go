package queue

// Queue 队列
type Queue[T any] struct {
	data  []T
	front int
}

// EmptyInfo 为空时的错误提示信息
const EmptyInfo = "queue is empty"

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
