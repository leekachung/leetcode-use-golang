func subsets(nums []int) [][]int {
    res := make([][]int, 0)
    helper(&res, nums, []int{}, 0)
    return res
}

func helper(res *[][]int, nums, temp []int, i int) {
    var e = make([]int, len(temp))
    copy(e, temp)
    *res = append(*res, e)
    for j := i; j < len(nums); j++ {
        temp = append(temp, nums[j])
        helper(res, nums, temp, j + 1)
        temp = temp[:len(temp)-1]
    }
}
