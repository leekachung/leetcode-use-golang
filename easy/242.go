func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    hash := make(map[rune]int)
    for _, v := range s {
        hash[v]++
    }
    for _, v := range t {
        if hash[v] != 0 {
            hash[v]--
        } else {
            return false
        }
    }
    return true
}
