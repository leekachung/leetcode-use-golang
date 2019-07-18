// solution 1 use func strings.Index()
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 { return "" }
    prefix := strs[0]
    for _, v := range strs {
        // 当strings.Index(v, prefix)不为0 则前缀不符合条件
	for strings.Index(v, prefix) != 0 {
            if len(prefix) == 0 { return "" }
            prefix = prefix[0:len(prefix)-1]
        }
    }
    return string(prefix)
}

// solution 2
func longestCommonPrefix(strs []string) string {
    length := len(strs)
    if length == 0 { return "" }
    minStr := strs[0]
    
    // 获取最短字符串
    for _, str := range strs {
        if len(str) < len(minStr) {
            minStr = str
        }
    }
    // 比较字符串的前缀
    for i, v := range minStr {
        for j := 0; j < length; j++ {
            // 当前缀不同 修改数组  需转换类型
            if byte(v) != strs[j][i] {
                return strs[j][:i]
            }
        }
    }
    return minStr
} 
