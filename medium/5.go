// solution 1 4ms
func longestPalindrome(s string) string {
    length := len(s)
    // 记录最长回文子串位置
    start, end := 0, 0
    
    for i := 0; i < length; i++ {
        // 中心扩散的"指针"
        left, right := i, i
        // 中心扩散
        for left >= 0 && right < length && s[left] == s[right] {
            left--
            right++
        }
        if right - left - 1 > end - start {
            start, end = left + 1, right
        }
        
        // 中心扩散的"指针"
        left, right = i, i + 1
        // 中心扩散
        for left >= 0 && right < length && s[left] == s[right] {
            left--
            right++
        }
        if right - left - 1 > end - start {
            start, end = left + 1, right
        }
    
        
    }
    
    return s[start:end]
}

// solution 2 8ms
func longestPalindrome(s string) string {
    var max string
    for i := 0; i < len(s); i++ {
        max = maxPalindrome(s, i, i, max)
        max = maxPalindrome(s, i, i+1, max)
    }
    return max
}

func maxPalindrome(s string, i, j int, max string) string {
    var sub string
    for i >= 0 && j < len(s) && s[i] == s[j] {
        sub = s[i:j+1]
        i--
        j++
    }
    if len(max) < len(sub) {
        return sub
    }
    return max
}

// solution 3 0ms
func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }
    begin, maxLen:=0, 1
    for i := 0 ; i< len(s);{
        if len(s) - i <= maxLen /2 {
            break
        }
        b, e := i, i
		for e < len(s) - 1 && s[e + 1] == s[e] {
			e++
        }
        i = e + 1
        for e < len(s) - 1 && b > 0 && s[e + 1] == s[b - 1] {
			e++
			b--
        }
        newLen := e + 1 - b
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}
    }
 
    return s[begin:begin + maxLen]
} 
