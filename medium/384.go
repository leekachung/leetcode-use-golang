type Solution struct {
    nums []int
}


func Constructor(nums []int) Solution {
    rand.Seed(time.Now().UnixNano())
    return Solution{ nums }
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	tmp := make([]int, len(this.nums))
	copy(tmp, this.nums)

	// Fisher–Yates shuffle 洗牌算法
	for i := len(tmp) - 1; i > 0; i-- {
		r := rand.Intn(i + 1)
		tmp[i], tmp[r] = tmp[r], tmp[i]
	}

	return tmp
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
