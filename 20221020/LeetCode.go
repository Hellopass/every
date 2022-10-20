/*
1003. 检查替换后的词是否有效 难度：中等
给你一个字符串 s ，请你判断它是否 有效 。
字符串 s 有效 需要满足：假设开始有一个空字符串 t = "" ，你可以执行 任意次 下述操作将 t 转换为 s ：

将字符串 "abc" 插入到 t 中的任意位置。形式上，t 变为 tleft + "abc" + tright，其中 t == tleft + tright 。注意，tleft 和 tright 可能为 空 。
如果字符串 s 有效，则返回 true；否则，返回 false。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/check-if-word-is-valid-after-substitutions
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//题解，两种解决方式
//第一种
func isValid(s string) bool {
  	var TimesString []rune
	a := 0
	for k := 0; k < len(s); k++ {
		if k+2 < len(s) && s[k] == 'a' && s[k+1] == 'b' && s[k+2] == 'c' {
			a++
			k += 2
		} else {
			TimesString = append(TimesString, rune(s[k]))
		}
	}

	if len(TimesString) == 0 {
		return true
	}

	if   a == 0 {
		return false
	}

	return isValid(string(TimesString))
 
	
}

//第二种

func isValid(s string) bool {
 for s != strings.ReplaceAll(s, "abc", "") {
		s = strings.ReplaceAll(s, "abc", "")
	}
	return s == ""
 
	
}




/*

分析
第一种方式相当于在实现第二种方法的底层代码，其中第二种解法strings.ReplaceAll()的主要功能就是将"abc"替换成""(空)
核心思想都是先将s中的"abc"给替换掉，在剩余的组成一个新的字符串在进行比较，第一种就是利用递归的方式将字符串逐一比较，第二种则直接调用库函数，在实际使用中建议使用第二种（出错率小）
*/
