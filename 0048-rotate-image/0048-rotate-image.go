func rotate(matrix [][]int)  {
    m:=len(matrix)
    n:=len(matrix[0])

    for i:=0;i<m;i++{
        for j:=i;j<n;j++{
            t:=matrix[i][j]
            matrix[i][j]=matrix[j][i]
            matrix[j][i]=t
        }
    }

    for i:=0;i<m;i++{
        slices.Reverse(matrix[i])
    }
}