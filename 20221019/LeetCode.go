/*
2279. 装满石头的背包的最大数量   难度：中等
现有编号从 0 到 n - 1 的 n 个背包。给你两个下标从 0 开始的整数数组 capacity 和 rocks 。第 i 个背包最大可以装 capacity[i] 块石头，当前已经装了 rocks[i] 块石头。另给你一个整数 additionalRocks ，表示你可以放置的额外石头数量，石头可以往 任意 背包中放置。

请你将额外的石头放入一些背包中，并返回放置后装满石头的背包的 最大 数量。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-bags-with-full-capacity-of-rocks
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

//题解
func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
//贪心
nums :=0
for i:=0;i<len(capacity);i++{
  capacity[i]-=rocks[i]
}
sort.Ints(capacity)
for j:=0;j<len(capacity);j++{
     if capacity[j]>=0&&(additionalRocks-capacity[j])>=0{
     additionalRocks-=capacity[j]
     nums++
}
}
      return nums
}

/*
分析 
1,由题目的意思可以知道，用addtionalRocks 来满足将背包装满的最大背包数
2,由1就可以知道我们需要从还需要石头最少的背包开始，也就是使用贪心的方法。
3,计算每个背包还差的石头数量，将这样的一个切片按照从小到大的顺序排列，从最小的开始放，直到addtionalRocks-capacity[i]<0,每一次nums++
*/
