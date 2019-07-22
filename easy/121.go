func maxProfit(prices []int) int {
    // 记录最小值
    min := 1 << 31 - 1
    // 记录最大利润
    max := 0
    for i := 0 ; i < len(prices); i++ {
        if min > prices[i] {
            min = prices[i]
        }
        if prices[i] - min > max {
            max = prices[i] - min
        }
    } 
    return max
}
