// solution 1
func hammingDistance(x int, y int) int {
    return bits.OnesCount(uint(x) ^ uint(y))
}

// solution 2
func hammingDistance(x int, y int) int {
    x = x ^ y
    res := 0
    for x != 0 {
        res++
        x &= x - 1
    }
    return res
}
