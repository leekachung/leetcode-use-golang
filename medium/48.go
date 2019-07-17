// Solution 1
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

// Solution 2
// 分别将四个顶点顺时针旋转
func rotate(matrix [][]int)  {
    n := len(matrix)
	len := n-1
	for i := 0; i < n/2; i++ {
		for j := i; j < len-i; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[len-j][i]
			matrix[len-j][i] = matrix[len-i][len-j]
			matrix[len-i][len-j] = matrix[j][len-i]
			matrix[j][len-i]  = tmp
		}
	}
}
