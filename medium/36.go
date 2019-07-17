func isValidSudoku(board [][]byte) bool {
    var rows, cols, boxs []map[byte]bool
    // 创建hash表 用于匹配是否出现过
    for i := 0; i <= 8; i++ {
        rows = append(rows, make(map[byte]bool))
        cols = append(cols, make(map[byte]bool))
        boxs = append(boxs, make(map[byte]bool))
    }
    
    // 遍历数独
    for i := range board {
        for j, v := range board[i] {
            if v == '.' {
                continue
            }
            // 重复出现
            if rows[i][v] || cols[j][v] || boxs[(i/3)*3+j/3][v] {
                return false
            }
            // 第一次出现
            rows[i][v] = true
            cols[j][v] = true
            boxs[(i/3)*3+j/3][v] = true
        }
    }
    return true
}
