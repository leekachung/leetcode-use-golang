func lengthOfLongestSubstring(s string) int {
    // 建立哈希表
    hash := make(map[byte]int)
    res := 0
    
    for start, end := 0, 0; end < len(s); end++ {
        if _, ok := hash[s[end]]; ok {
            start = int(math.Max(float64(hash[s[end]]), float64(start)))
        }
        res = int(math.Max(float64(res), float64(end - start + 1)))
        hash[s[end]] = end + 1
    }
    
    return res
}
