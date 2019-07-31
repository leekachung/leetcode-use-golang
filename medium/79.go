func exist(board [][]byte, word string) bool {
    if len(word) == 0 { return false }
    b := []byte(word)
    for x := range board {
        for y := range board[x] {
            if dfs(x, y, board, b) {
                return true
            }
        }
    }
    return false
}

func dfs(x, y int, board [][]byte, word []byte) bool {
    if x < 0 || x >= len(board) { return false }
    if y < 0 || y >= len(board[x]) { return false }
    if board[x][y] != word[0] { return false }
    if len(word) == 1 { return true }
    temp := board[x][y]
    board[x][y] = '-'
    found := dfs(x + 1, y, board, word[1:]) ||
        dfs(x, y + 1, board, word[1:]) ||
        dfs(x - 1, y, board, word[1:]) ||
        dfs(x, y - 1, board, word[1:])
    board[x][y] = temp
    
    return found
}
