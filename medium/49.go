// solution 1 有溢出的风险
func groupAnagrams(strs []string) [][]string {
    // 累计字母异位词的个数
    temp := 0
    // 建立哈希表
    hash := make(map[int]int)
    // 用26个质数代表26个字母
    prime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}
    // 结果集
    res := [][]string{}
    
    for _, s := range strs {
        p := 1
        for _, ss := range s {
            // 计算字母异位词的值
            p *= prime[ss - 'a']
        }
        // 首次出现的字母异位词
        if _, ok := hash[p]; !ok {
            hash[p] = temp
            temp++
            res = append(res, []string{})
        }
        res[hash[p]] = append(res[hash[p]], s)
    }

    return res
}

// solution 2 不会有溢出的风险 原理和 solution 1 相似
func groupAnagrams(strs []string) [][]string {
    // 累计字母异位词的个数
    temp := 0
    // 建立哈希表
    hash := make(map[string]int)
    // 结果集
    res := [][]string{}
    
    for _, s := range strs {
        nums := make([]int, 26)
        for _, ss := range s {
            nums[ss - 'a']++
        }
        key := ""
        for j := 0; j < len(nums); j++ {
            key = key + string(nums[j]) + "#";
        }
        if _, ok := hash[key]; !ok {
            hash[key] = temp
            temp++
            res = append(res, []string{})
        }
        res[hash[key]] = append(res[hash[key]], s)
    }
  
    return res
}

// solution 3 先排序字符串的字符顺序 再哈希匹配
func groupAnagrams(strs []string) [][]string {
    m := map[string]int{}
    ret := [][]string{}
    for _, v := range strs {
        arr := strings.Split(v, "")
        sort.Strings(arr)
        vv := strings.Join(arr, "")
        if i, ok := m[vv]; ok {
            ret[i] = append(ret[i], v)
        } else {
            m[vv] = len(ret)
            ret = append(ret, []string{v})
        }
    }
    return ret
}
