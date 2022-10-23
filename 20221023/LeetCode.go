/*
118. 杨辉三角 难度简单
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。

*/



//解法
func generate(numRows int) [][]int {
    ans := make([][]int, numRows)
    for i := range ans {
        ans[i] = make([]int, i+1)
        ans[i][0] = 1
        ans[i][i] = 1
        for j := 1; j < i; j++ {
            ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
        }
    }
    return ans
}


/*分析:

1,主要知道杨辉三角的特点就是等于上一行的两个相邻数相加
2，边界都是1
3,首先设置边界，之后根据定义计算ans[i-1][j] + ans[i-1][j-1]
*/
