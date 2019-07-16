func maxProfit(prices []int) int {
    // 数组为空 直接返回
    if len(prices) == 0 {
        return 0
    }
    // 利润
    temp := 0
    for i := 0; i < len(prices) - 1; i++ {
        // 当出现第二天价格高于当天 则当天买进 第二天卖出
        if ( prices[i] < prices[i+1] ) {
            temp += prices[i+1] - prices[i]
        }
    }
    return temp
}
