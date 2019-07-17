func reverse(x int) int {
    res := 0
    for x != 0 {
        res = res * 10 + x % 10
        if res < math.MinInt32 || res > math.MaxInt32 {
            return 0
        }
        x /= 10
    }
    return res
}
