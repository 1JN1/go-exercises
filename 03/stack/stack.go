package stack

// Stack 栈
type Stack[T any] struct {
	data []T
	top  int
}

// DefaultCapacity 初始容量
const DefaultCapacity = 10

// EmptyInfo 栈为空时的错误提示信息
const EmptyInfo = "stack is empty"

// NewStack 构造函数
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{make([]T, DefaultCapacity), -1}
}

// Push 入栈
func (s *Stack[T]) Push(v T) {

	s.top++

	s.data[s.top] = v
}

// Pop 出栈
func (s *Stack[T]) Pop() T {

	if s.IsEmpty() {
		panic(EmptyInfo)
	}

	ret := s.data[s.top]

	s.top--

	return ret
}

// Peek 获取栈顶元素
func (s *Stack[T]) Peek() T {

	if s.IsEmpty() {
		panic(EmptyInfo)
	}

	return s.data[s.top]
}

// IsEmpty 判断栈是否为空
func (s *Stack[T]) IsEmpty() bool {

	return s.top == -1
}
