/*567. 字符串的排列 难度 中等
经典的双指针问题
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。

换句话说，s1 的排列之一是 s2 的 子串 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/permutation-in-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

示例 1：

输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").

示例 2：

输入：s1= "ab" s2 = "eidboaoo"
输出：false

*/

//题解
func checkInclusion(s1, s2 string) bool {
    n, m := len(s1), len(s2)
    if n > m {
        return false
    }
    var cnt1, cnt2 [26]int
    for i, ch := range s1 {
        cnt1[ch-'a']++
        cnt2[s2[i]-'a']++
    }
    if cnt1 == cnt2 {
        return true
    }
    for i := n; i < m; i++ {
        cnt2[s2[i]-'a']++
        cnt2[s2[i-n]-'a']--
        if cnt1 == cnt2 {
            return true
        }
    }
    return false
}
/*
分析总结
1，滑动窗口问题，一个固定的窗口问题
2，我们将窗口大小设置为s1的长度,其中第一次的遍历就是将cnts2的窗口大小设置为len(s1)
3,每一次都向后移动一位，将前面的字符减去，再将后面一位的字符加上，相当于窗口整体向右移动一位，直到每个字符出现先得次数等于所给出得字符次数，并且s1要小于s2  



滑动窗口问题也成为双指针问题，有定长窗口，有变长窗口，其中定长窗口就一步一步往下遍历，而变长窗口要注意题目所给条件，不断调整左右指针得指向
*/
