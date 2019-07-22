// solution 1 T(n) = O(n) S(n) = O(n)
func climbStairs(n int) int {
    if n == 1 { return 1 }
    count := make([]int, n + 1)
    count[1] = 1
    count[2] = 2
    for i := 3; i <= n; i++ {
        count[i] = count[i - 1] + count[i - 2]
    }
    return count[n]
}

// solution 2 T(n) = O(n) S(n) = O(1)
func climbStairs(n int) int {
    if n == 1 { return 1 }
    one := 1
    two := 2
    for i := 3; i <= n; i++ {
        three := one + two
        one = two
        two = three
    }
    return two
}

// solution 3 T(n) = O(log(n)) S(n) = O(1)
func climbStairs(n int) int {
	sqrt5 := math.Sqrt(5)
	fibN := math.Pow((1+sqrt5)/2, float64(n+1)) - math.Pow((1-sqrt5)/2, float64(n+1))
	return int(math.Round(fibN / sqrt5))
}
