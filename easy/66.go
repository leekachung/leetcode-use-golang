func plusOne(digits []int) []int {
    for i := len(digits) - 1; i >= 0; i-- {
        // 当前位小于9 不进位 加1返回
        if digits[i] < 9 {
            digits[i]++
            return digits
        } else {
            // 进位
            digits[i] = 0
        }
    }
    return append([]int{1}, digits...)
}
