func rob(nums []int) int {
    prev := 0
    curr := 0
    for _, v := range nums {
        temp := curr
        curr = int(math.Max(float64(prev + v), float64(curr)))
        prev = temp
    }
    return curr
}
