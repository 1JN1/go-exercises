package stack

// Stack 栈
type Stack[T any] struct {
	data []T
	top  int
}

// EmptyInfo 栈为空时的错误提示信息
const EmptyInfo = "stack is empty"

// Push 入栈
func (s *Stack[T]) Push(v T) {

	s.top++
	s.data = append(s.data, v)
}

// Pop 出栈
func (s *Stack[T]) Pop() T {

	if s.IsEmpty() {
		panic(EmptyInfo)
	}

	s.top--

	return s.data[s.top]
}

// Peek 获取栈顶元素
func (s *Stack[T]) Peek() T {

	if s.IsEmpty() {
		panic(EmptyInfo)
	}

	return s.data[s.top-1]
}

// IsEmpty 判断栈是否为空
func (s *Stack[T]) IsEmpty() bool {

	return s.top == 0
}
