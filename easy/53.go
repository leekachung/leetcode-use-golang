func maxSubArray(nums []int) int {
    sum := 0
    ans := nums[0]
    for _, v := range nums {
        if sum > 0 {
            sum += v
        } else {
            sum = v
        }
        ans = int(math.Max(float64(sum), float64(ans)))
    }
    return ans
}
