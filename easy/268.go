// solution 1 不适合数字重复出现的场景
func missingNumber(nums []int) int {
    sum := 0
    total := len(nums)
    for i := 0; i < len(nums); i++ {
        total += i
        sum += nums[i]
    }
    return total - sum
}

// solution 2 适合数字随机重复出现的场景 
func missingNumber(nums []int) int {
    m := make([]byte, len(nums) + 1)
    for _, v := range nums {
        m[v] = 1
    }
    for k, v := range m {
        if v < 1 {
            return k
        }
    }
    return -1    
}

// solution 3 适合数字以奇数次重复出现的场景
func missingNumber(nums []int) int {
    res := len(nums)
    for i := 0; i < len(nums); i++ {
        res ^= i ^ nums[i]
    }
    return res
}

