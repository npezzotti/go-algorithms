package implement_queue_using_stacks

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() (x int) {
	x = s.Peek()
	*s = (*s)[:len(*s)-1]
	return
}

func (s *Stack) Peek() (x int) {
	x = (*s)[len(*s)-1]
	return
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

type MyQueue struct {
	pushStack Stack
	popStack  Stack
}

func Constructor() MyQueue {
	return MyQueue{
		pushStack: make(Stack, 0),
		popStack:  make(Stack, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.pushStack = append(this.pushStack, x)
}

func (this *MyQueue) Pop() int {
	if this.popStack.Empty() {
		for len(this.pushStack) > 0 {
			this.popStack.Push(this.pushStack[len(this.pushStack)-1])
			this.pushStack = this.pushStack[:len(this.pushStack)-1]
		}
	}

	return this.popStack.Pop()
}

func (this *MyQueue) Peek() int {
	if this.popStack.Empty() {
		for len(this.pushStack) > 0 {
			this.popStack.Push(this.pushStack[len(this.pushStack)-1])
			this.pushStack = this.pushStack[:len(this.pushStack)-1]
		}
	}

	return this.popStack.Peek()
}

func (this *MyQueue) Empty() bool {
	return len(this.pushStack) == 0 && len(this.popStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
