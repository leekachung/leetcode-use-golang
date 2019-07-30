// solution 1
func letterCombinations(digits string) []string {
    res := []string{}
    if len(digits) == 0 { return res }
    
    // 创建匹配规则
    rules := map[rune]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno",
		'7': "pqrs", '8': "tuv", '9': "wxyz",
	}
    
    for i, v := range digits {
        if i == 0 {
            for _, r := range rules[v] {
                res = append(res, string(r))
            }
        } else {
            temp := res
            res = nil
            for _, v1 := range temp {
				for _, v2 := range rules[v] {
					res = append(res, v1+string(v2))
				}
			}
        }
    }
    return res
}
