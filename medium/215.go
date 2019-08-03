// solution 1
func findKthLargest(nums []int, k int) int {
    sort.Ints(nums)
    i := len(nums) - 1
    for i > 0 && k > 1 {
        k--
        i--
    }
    return nums[i]
}

// solution 2
func maxHeapify(nums []int, start, end int) {
	var father = start
	var child = 2 * father + 1

	for child <= end {
		if child + 1 <= end && nums[child] < nums[child + 1] {
			child++
		}
		if nums[father] > nums[child] {
			return
		} else {
			nums[father], nums[child] = nums[child], nums[father]
			father = child
			child = 2 * father + 1
		}
	}
}

func heapSort(nums []int, k int) {
	for i := len(nums) / 2 -1; i >=0 ; i-- {
		maxHeapify(nums, i, len(nums) - 1)
	}
	for i := len(nums) - 1; i >= len(nums) - k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(nums, 0, i - 1)
	}
}


func findKthLargest(nums []int, k int) int {
	heapSort(nums, k)
	return nums[len(nums) - k]
}
