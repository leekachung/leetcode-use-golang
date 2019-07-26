func isValid(s string) bool {
    stack := make([]rune, 0)
    
    // 括号映射表
    rightMap := map[rune]rune{ ')':'(', ']':'[', '}':'{' }
    
    for _, v := range s {
        if v == '(' || v == '{' || v == '[' {
            stack = append(stack, v)
        } else if v == ')' || v == '}' || v == ']' {
            if len(stack) != 0 && stack[len(stack)-1] == rightMap[v] {
                stack = stack[0:len(stack)-1]
            } else {
                return false
            }
        }
    }
    
    if len(stack) == 0 {
        return true
    } else {
        return false
    }
}
