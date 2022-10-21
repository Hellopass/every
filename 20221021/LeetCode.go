/*
88. 合并两个有序数组  难度：简单
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/merge-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/



//解法
//第一种
func merge(nums1 []int, m int, nums2 []int, n int)  {
    j :=0
    for i:=m;i<m+n;i++{ 
     if j<n{
    nums1[i]=nums2[j]
     j++
}
    }
//   copy(nums1[m:], nums2)
   sort.Ints(nums1)
}
此时的时间复杂度为O(n)


//第二种
func merge(nums1 []int, m int, nums2 []int, n int)  {
//     j :=0
//     for i:=m;i<m+n;i++{ 
//      if j<n{
//     nums1[i]=nums2[j]
//      j++
// }
//     }
  copy(nums1[m:], nums2)
   sort.Ints(nums1)
}
此时时间复杂度为O(n^2)

/*
分析:
1,首先题目要求合并到nums1，我们就找到nums1后面的空位将nums2插入即可,或者使用copy(type another ...)来进行覆盖
*/

