func generate(numRows int) [][]int {
    res := [][]int{[]int{1}, []int{1, 1}}
    for i:= 2; i < numRows; i++ {
        temp := []int{1}
        for j := 1; j < i; j++ {
            temp = append(temp, res[i-1][j-1] + res[i-1][j])
        }
        temp = append(temp, 1)
        res = append(res, temp)
    }
    return res[:numRows]
}
