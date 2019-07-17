func reverseString(s []byte)  {
    length := len(s) - 1
    for i, _ := range s {
        if i > length/2 {
            break
        }
        s[i], s[length-i] = s[length-i], s[i]
    }
}
