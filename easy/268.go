func missingNumber(nums []int) int {
    sum := 0
    total := len(nums)
    for i := 0; i < len(nums); i++ {
        total += i
        sum += nums[i]
    }
    return total - sum
}
