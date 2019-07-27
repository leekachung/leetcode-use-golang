// solution 1 32ms
func setZeroes(matrix [][]int)  {
    if len(matrix) == 0 { return }
    isCol := false
    big := len(matrix)
    small := len(matrix[0])
    // 定位 0 元素
    for i := 0; i < big; i++ {
        // 首行包含 0 元素
        if matrix[i][0] == 0 {
            isCol = true
        }
        for j := 1; j < small; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }
    // 除首行首列 置零
    for i := 1; i < big; i++ {
        for j := 1; j < small; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }    
    // 首个元素是否为零 将首列置零
    if matrix[0][0] == 0 {
        for j, _ := range matrix[0] {
            matrix[0][j] = 0
        }
    }
    // 将首行置零
    if isCol {
        for i, _ := range matrix {
            matrix[i][0] = 0
        }
    }
    
}

// solution 2 24ms
func setZeroes(matrix [][]int)  {
    if len(matrix)==0{return}
    n,m:=len(matrix),len(matrix[0])
    flag := make([]bool, n+m,n+m)
    for i:=0;i<n;i++{
        for j:=0;j<m;j++{
            if matrix[i][j]==0{
                flag[i]=true
                flag[n+j]=true
            }
        }
    }
    for i:=0;i<n;i++{
        for j:=0;j<m;j++{
            if flag[i]||flag[n+j] {matrix[i][j]=0}
        }
    }
}
