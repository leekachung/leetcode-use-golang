func removeDuplicates(nums []int) int {
    // 数组为空
    if len(nums) == 0 {
        return 0
    }
    // 计算重复值的偏移量
    size := 0
    for i,_ := range nums {
        if nums[i] != nums[size] {
            size += 1
            nums[size] = nums[i]
        }
    }
    // go切片特性 修改数组元素
    nums = nums[0:size+1]
    return size + 1
}
