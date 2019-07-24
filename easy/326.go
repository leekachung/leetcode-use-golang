func isPowerOfThree(n int) bool {
    if n <= 0 { return false }
    for n % 3 == 0 {
        n = n / 3
    }
    return 1 == n
}
