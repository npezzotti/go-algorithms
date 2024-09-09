package min_stack

import "math"

type MinStack struct {
	stack []int
	min   int
}

func Constructor() MinStack {
	return MinStack{
		min: math.MaxInt64,
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if this.min > val {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	last := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	if last == this.min {
		this.min = math.MaxInt64
		for _ ,x := range this.stack {
			if x < this.min {
				this.min = x
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
