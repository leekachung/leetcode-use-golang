func reverse (nums []int) {
    i, j := 0, len(nums) - 1
    for i < j {
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
}

func rotate(nums []int, k int)  {
    k %= len(nums)
    reverse(nums[:len(nums)-k])
    reverse(nums[len(nums)-k:])
    reverse(nums)
}
