/*
73. 矩阵置零  难度：中等
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
*/

//解法
func setZeroes(matrix [][]int)  {
nums :=matrix
i,j:=0,0
var v []int
Mapx :=make(map[int]bool)
Mapy :=make(map[int]bool)
for i,v=range matrix{
    for j=range v{
        if matrix[i][j]==0{
            Mapx[i]=true
            Mapy[j]=true
        }
    }
}
for i, r := range matrix {
        for j := range r {
            if Mapx[i] || Mapy[j] {
                r[j] = 0
            }
        }
    }



matrix =nums
}

/*
分析
1，我们遍历一遍数组找到所对应对的位置，将相对位置存在不同的map中，假设为x,y，则将x,y存入map中
2，根据上面的我们的到了0的位置，然后0所在在列行都是0，我们只需要遍历数组找出那个数的x,y是属于我们找到的x,y，再将他复制为0
*/
