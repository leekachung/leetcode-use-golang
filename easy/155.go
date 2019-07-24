// solution 1 28ms
type MinStack struct {
    nums []int
    min  []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(x int)  {
    this.nums = append(this.nums, x)
    if len(this.min) == 0 || x <= this.min[len(this.min)-1] {
        this.min = append(this.min, x)
    }
}
    

func (this *MinStack) Pop()  {
    if len(this.nums) == 0 { return }
    // 即将pop的值
    temp := this.nums[len(this.nums)-1]
    this.nums = this.nums[0:len(this.nums)-1]
    if len(this.min) != 0 && temp == this.min[len(this.min)-1] {
        this.min = this.min[0:len(this.min)-1]
    }
}


func (this *MinStack) Top() int {
    if len(this.nums) == 0 { return 0 }
    return this.nums[len(this.nums)-1]
}


func (this *MinStack) GetMin() int {
    if len(this.min) == 0 { return 0 }
    return this.min[len(this.min)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// solution 2 20ms
type MinStack struct {
    Xstack []int // 存放push到stack的x
    Mstack []int // 存放min(stack[-1], x)
}


/** initialize your data structure here. */
func Constructor() MinStack {
    
    return MinStack{[]int{}, []int{math.MaxInt32}}
}


func (this *MinStack) Push(x int)  {
    this.Xstack = append(this.Xstack, x)
    this.Mstack = append(this.Mstack, min(x, this.Mstack[len(this.Mstack)-1]))
}


func (this *MinStack) Pop()  {
    this.Xstack = this.Xstack[:len(this.Xstack)-1]
    this.Mstack = this.Mstack[:len(this.Mstack)-1]
}


func (this *MinStack) Top() int {
    return this.Xstack[len(this.Xstack)-1]
}


func (this *MinStack) GetMin() int {
    return this.Mstack[len(this.Mstack)-1]
}

func min(a, b int)int{
    if a < b{
        return a
    }
    return b
}
/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
