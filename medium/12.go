func intToRoman(num int) string {
    res := ""
    if num < 1 || num > 3999 { return res }
    // 转化匹配规则
    values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	letters := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    for i := 0; i < len(values); {
        if num >= values[i] {
            num -= values[i]
            res += letters[i]
        } else {
            // 只有当values[i]的值大于num当前值 i自增
            i++
        }
    }
    return res
}
