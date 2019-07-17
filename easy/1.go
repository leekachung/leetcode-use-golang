// Solution 1
func twoSum(nums []int, target int) []int {
    hash := make(map[int]int)
    res := make([]int, 2)
    for i, v := range nums {
        hash[v] = i
    }
    for ii, vv := range nums {
        res[0] = ii
        temp := target - vv
        index, ok := hash[temp]
        if index != ii && ok {
            res[1] = index
            break
        }
    }
    return res
}

// Solution 2  better
func twoSum(nums []int, target int) []int {
    hash := make(map[int]int)
    res := make([]int, 2)
    for i, v := range nums {
        temp := target - v
        ii, ok := hash[temp]
        if ok {
            res[0] = i
            res[1] = ii
        }
        hash[v] = i
    }
    return res
}
