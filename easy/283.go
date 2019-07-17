func moveZeroes(nums []int)  {
    i := 0
    length := len(nums)
    for _, v := range nums {
        if v != 0 {
            nums[i] = v
            i++
        }
    }
    for j := i; j < length; j++ {
        nums[j] = 0
    }
}
