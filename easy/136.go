// Solution 1 T(n) = O(n) | S(n) = O(n)
func singleNumber(nums []int) int {
    hash := make(map[int]bool)
    for _, v := range nums {
        if _, ok := hash[v]; ok {
            hash[v] = false
        } else {
            hash[v] = true
        }
    }
    for i, val := range hash {
        if val {
            return i
        }
    }
    return 0
}

// Solution 2 T(n) = O(n) | S(n) = S(1)
func singleNumber(nums []int) int {
    result := 0
    for _, v := range nums {
        result = result ^ v
    }
    return result
}
