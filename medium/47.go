func permuteUnique(nums []int) [][]int {
    res := make([][]int, 0)
    sort.Ints(nums)
    backtrack(nums, []int{}, &res)
    return res
}

func backtrack(nums, prev []int, res *[][]int) {
    if len(nums) == 0 { 
        *res = append(*res, append([]int{}, prev...))
        return
    }
    
    for i, v := range nums {
        if i != 0 && nums[i] == nums[i-1] {
			continue
		}
        temp := append(append([]int{}, nums[0:i]...), nums[i+1:]...)
        backtrack(temp, append(prev, v), res)
    }
}
