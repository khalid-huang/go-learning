//策略是定下终点，然后往前去找最远的点，但是有一点要注意就是在往前找的时间，查找的范围要确保是无相同的
package main

//可以用动规，可以用暴力查找

//暴力查找，定好当前位置，然后往前找最远的，不过要注意查找的范围无相同(（i-1)-c[i-1]+1）
//动规：https://blog.csdn.net/qq_38494372/article/details/80846954
//用left来跟踪到目前为止，没有出现重复段的最左边的字符，然后当出现重复字符的时候就判断下left与上一个重复字条之间的大小；如果上一个重复字符的位置在left的右边，就需要更新left；否则判断加上当前字符是否出现更大的长度值
func lengthOfLongestSubstring(s string) int {
	location := [256]int{} //都是ASCII码

	for i := range location {
		location[i] = -1 //表示都没有见过
	}

	maxLen, left := 0, 0
	for i := 0; i < len(s); i++ {
		//s[i]上次出现的位置在左边没有重复字段最左边位置的右边，也就是没有重复字段中有s[i]
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i+1-left > maxLen {
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}
	return maxLen
}

func test_3() {
}
