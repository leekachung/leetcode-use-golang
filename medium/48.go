func rotate(matrix [][]int)  {
    // 对角线翻转
    length := len(matrix)
    for i := 0; i < length-1; i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
    // 水平翻转
    for i := 0; i < length/2; i++ {
        for j := range matrix[i] {
            matrix[j][i], matrix[j][length-i-1] = matrix[j][length-i-1], matrix[j][i]
        }
    }
}
