package go_code

type MyStack struct {
	queue []int
}

func Constructor() (s MyStack) {
	return
}

func (s *MyStack) Push(x int) {
	s.queue = append(s.queue, x)
}

func (s *MyStack) Pop() int {
	res := s.queue[len(s.queue)-1]
	s.queue = s.queue[:len(s.queue)-1]
	return res
}

func (s *MyStack) Top() int {
	return s.queue[len(s.queue)-1]
}

func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
