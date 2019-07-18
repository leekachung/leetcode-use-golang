const (
    // 允许的返回值范围
    maxInt = 1 << 31 -1
    minInt = -1 << 31
)

func myAtoi(str string) int {
    length := len(str)
    var i, j, res int
    // 记录正负符号 true为正 false为负
    flag := true
    
    if length == 0 { return 0 }
    // 定位数字开头
    for i = 0; i < length; i++ {
        if str[i] == '+' {
            i++
            break 
        }  else if str[i] == '-' { 
            i++
            flag = false 
            break
        } else if str[i] >= '0' && str[i] <= '9' { 
            break 
        } else if str[i] != ' ' { 
            return 0 
        }
    }
    // 定位数字结尾
    for j = i; j < length; j++ {
        if !(str[j] >= '0' && str[j] <= '9') { 
            break
        }
    }
    // 组合数字
    for k := i; k < j; k++ {
        num := int(str[k] - '0')
        if flag {
            res = res*10 + num
            if res > maxInt {
                return maxInt
            }
        } else {
            res = res*10 - num
            if res < minInt {
                return minInt
            }
        }
    }
    return res
}
