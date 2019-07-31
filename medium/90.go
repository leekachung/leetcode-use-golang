func subsetsWithDup(nums []int) [][]int {
    res := make([][]int, 0)
    sort.Ints(nums)
    helper(&res, nums, []int{}, 0)
    return res
}

func helper(res *[][]int, nums, temp []int, i int) {
    var e = make([]int, len(temp))
    copy(e, temp)
    *res = append(*res, e)
    for j := i; j < len(nums); j++ {
        //和上个数字相等就跳过
        if (j > i && nums[j] == nums[j - 1]) {
            continue;
        }

        temp = append(temp, nums[j])
        helper(res, nums, temp, j + 1)
        temp = temp[:len(temp)-1]
    }
}
