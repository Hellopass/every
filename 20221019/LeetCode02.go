/*
299. 猜数字游戏 难度：中等
你在和朋友一起玩 猜数字（Bulls and Cows）游戏，该游戏规则如下：

写出一个秘密数字，并请朋友猜这个数字是多少。朋友每猜测一次，你就会给他一个包含下述信息的提示：

猜测数字中有多少位属于数字和确切位置都猜对了（称为 "Bulls"，公牛），
有多少位属于数字猜对了但是位置不对（称为 "Cows"，奶牛）。也就是说，这次猜测中有多少位非公牛数字可以通过重新排列转换成公牛数字。
给你一个秘密数字 secret 和朋友猜测的数字 guess ，请你返回对朋友这次猜测的提示。

提示的格式为 "xAyB" ，x 是公牛个数， y 是奶牛个数，A 表示公牛，B 表示奶牛。

请注意秘密数字和朋友猜测的数字都可能含有重复数字。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/bulls-and-cows
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

//题解
func getHint(secret string, guess string) string {
	bulls := 0
	cows := 0
	Nsecret := make(map[rune]int)
	Nguess := make(map[rune]int)

	for k, v := range secret {
		if v == rune(guess[k]) {
			bulls++
		} else {
			Nsecret[v]++
			Nguess[rune(guess[k])]++
		}
	}
	for k1, v1 := range Nsecret {
		if v1 < Nguess[k1] {
			cows += v1

		} else {
			cows += Nguess[k1]
		}
	}
	e := strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
	return e
}
/*
分析 ：
1,首先找到公牛,利用map使得string中得每一个数字得到储存map[rune]int 列如map[1]2表示数字1出现了2次
2，之后遍历得到得奶牛map，找到每个数次出现得次数，并选取最小得次数进行返回。
*/
