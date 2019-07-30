func permute(nums []int) [][]int {
    res := make([][]int, 0)
    backtrack(nums, []int{}, &res)
    return res
}

func backtrack(nums, prev []int, res *[][]int) {
    if len(nums) == 0 { 
        *res = append(*res, prev)
        return
    }
    
    for i, v := range nums {
        temp := append(append([]int{}, nums[:i]...), nums[i+1:]...)
        backtrack(temp, append(prev, v), res)
    }
}
