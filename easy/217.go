func containsDuplicate(nums []int) bool {
    // 声明hash表
    hash := make(map[int]bool)
    for _, v := range nums {
        // 判断hash表是否存在键值
        if _, ok := hash[v]; ok {
            return true
        }
        // 出现过的值置于true
        hash[v] = true
    }
    return false
}
