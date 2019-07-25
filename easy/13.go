// solution 1
func romanToInt(s string) int {
    // 字符数字映射表
    romaNums := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    
    // 特殊规则
    romaSpecial := map[byte][]byte {
        'I': {'V', 'X'},
        'X': {'L', 'C'},
        'C': {'D', 'M'},
    }
    
    sum := 0
    for i := 0; i < len(s); i++ {
        // 获取当前值
        curr := s[i]
        num := romaNums[curr]
        // 获取下一个值 若存在
        var next byte
        if i + 1 <= len(s) - 1 {
            next = s[i + 1]
        }
        if ss, ok := romaSpecial[curr]; ok {
            // 匹配到满足特俗规则
            for _, v := range ss {
                if next == v {
                    num = romaNums[next] - num
                    // 跳过下一个值 防止重复
                    i++
                    break
                }
            }
        }
        sum += num
    }
    return sum
}

// solution 2 less code
var romanInts = map[byte]int{'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}

func romanToInt(s string) int {
    prev := 3000
    ans := 0
    for i := 0; i < len(s); i++ {
        n := romanInts[s[i]]
        if n > prev {
            ans += n - 2*prev
        } else {
            ans += n
        }
        prev = n 
    }
    return ans
}
