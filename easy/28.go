func strStr(haystack string, needle string) int {
    n_length := len(needle)
    h_length := len(haystack)
    if n_length == 0 { return 0 }
    for i, _ := range haystack {
        // 当needle长度大于haystack 以及定位后的第一个字母索引加上needle长度大于haystack 
        // 可判定包含关系不存在
        if n_length > h_length || i + n_length > h_length{
            break
        }
        if haystack[i:i+n_length] == needle {
            return i
        }
    }
    return -1
}
