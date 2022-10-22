/*
121. 买卖股票的最佳时机*  难度简单
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

 

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
/


//解法
func maxProfit(prices []int) int {
    low := math.MaxInt32
    ans := 0

    for _, price := range prices{
        low = min(low, price)
        ans = max(ans, price-low)
    }
    return ans
}

func min(x, y int) int{
    if x<y{
        return x
    }
    return y
}

func max(x, y int) int{
    if x>y{
        return x
    }
    return y
}

/*分析

1,首先定义最小值等于一个极限值,min和max函数分别求两个数的最大值和最小值,
2,遍历价格找到最小值
3，ans每次保存所找到最低点于现在的差值并去上一次比较，保存每较大的ans
*/


