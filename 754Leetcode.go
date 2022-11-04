/*754. 到达终点数字

在一根无限长的数轴上，你站在0的位置。终点在target的位置。

你可以做一些数量的移动 numMoves :

每次你可以选择向左或向右移动。
第 i 次移动（从  i == 1 开始，到 i == numMoves ），在选择的方向上走 i 步。
给定整数 target ，返回 到达目标所需的 最小 移动次数(即最小 numMoves ) 。



示例 1:

输入: target = 2
输出: 3
解释:
第一次移动，从 0 到 1 。
第二次移动，从 1 到 -1 。
第三次移动，从 -1 到 2 。
示例 2:

输入: target = 3
输出: 2
解释:
第一次移动，从 0 到 1 。
第二次移动，从 1 到 3 。
*/

//题解
func reachNumber(target int) int {
	if target < 0 {
		target *= (-1)
	}
	nums := 0
	i := 0
	for nums < target || (nums-target)%2 != 0 {
		nums += i
		i++
	}
	return i-1
}
/*分析
1，这道题目是一道数学题
2，首先我们先考虑target的正负性，还有对称性，所以我们将target取绝对值
3,我们来看：1+2+3+4+5=15，1+2—3+4+5=9，我们发现15=9+3*2，所以得出来的结论是当一串自然数相加改变某个数在其中的符号，相当于给结果改变了2*A个值，就比如刚那个提到的，15就是9加上3的双倍
4，第一种情况是当相加恰好是target,第二种就是我们第三条说的，我们优先考虑相加，之后在相加之间改变符号，而且A*2必然是一个偶数，所以我们只需要判断这个和是不是小于nums或者差值是不是2的倍数
*/