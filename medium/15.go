func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := make([][]int, 0)
    if len(nums) < 3 { return res }
    
    // 表示首指针 代表第一个数
    for i := 0; i < len(nums); i++ {
        // 当第一个数大于0 则不可能和为0
        if nums[i] > 0 { break }
        // 去重
        if i > 0 && nums[i] == nums[i - 1] { continue } 
        // 表示尾指针 代表第三个数
        r := len(nums) - 1
        // 表示中指针 代表第二个数
        l := i + 1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            if sum == 0 {
                res = append(res, []int{nums[i], nums[l], nums[r]})
                for l < r && nums[l] == nums[l+1] { l++ }
                for l < r && nums[r] == nums[r-1] { r-- }
                l++
                r--
            } else if sum > 0 {
                r--
            } else if sum < 0 {
                l++
            }
        }
    }
    return res
}
