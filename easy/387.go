// Solution 1 use map
func firstUniqChar(s string) int {
    hash := make(map[rune]int)
    for _, v := range s {
        hash[v]++
    }
    for i, v := range s {
        if hash[v] == 1 {
            return i
        }
    }
    return -1
}

// Solution 2 use array
func firstUniqChar(s string) int {
    arr := make([]int, 26)
    for _, v := range s {
        arr[v - 'a']++
    }
    for i, v := range s {
        if arr[v - 'a'] == 1 {
            return i
        }
    }
    return -1
}
