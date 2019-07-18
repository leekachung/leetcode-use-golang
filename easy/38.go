// solution 1
func countAndSay(n int) string {
    // 递归终止条件
    if n == 1 {
        return "1"
    }
    // 递归读取上一个字符串
    s := countAndSay(n - 1)
    // 结果数组，偶数位为字符个数，奇数位字符本身
    nums := []int{}
    // 前一个字符
    var prevByte byte
    // 遍历字符串
    for i := 0; i <= len(s) - 1; i++ {
        if prevByte != s[i] { // 当前字符和前一个字符不同，则往结果数组中，插入当前字符个数1和字符本身
            nums = append(nums, 1, int(s[i] - '0'))
            prevByte = s[i]
        } else { // 当前字符和前一个字符相同，则更新当前字符的个数
            nums[len(nums) - 2]++
        }
    }
    return split(nums)
}
// 字符串数组拆分成string
func split(nums []int) string {
    result := []string{}
    for i := 0; i <= len(nums) - 1; i++ {
        result = append(result, strconv.Itoa(nums[i]))
    }
    return strings.Join(result, "")
}

// solution 2 less code
func countAndSay(n int) string {
   if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)
	r := make([]byte, 0)
	for i := 0; i < len(s); {
		j := i + 1
		for j < len(s) {
			if s[j] == s[i] {
				j++
				continue
			} else {
				break
			}
		}
		if j > i+1 {
			r = append(r, []byte(strconv.Itoa(j-i))...)
			r = append(r, s[i])
			i = j
		} else {
			r = append(r, '1', s[i])
			i++
		}
	}
	return string(r)
}
