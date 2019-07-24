func countPrimes(n int) int {
    count := 0
    temp := make([]int, n)
    for i := 2; i < n; i++ {
        if temp[i] == 1 { continue }
        for j := 2*i; j < n; j += i {
            temp[j] = 1
        }
        count++
    }
    return count
}
