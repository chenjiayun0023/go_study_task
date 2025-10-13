package main

import "fmt"

//https://leetcode.cn/problems/longest-common-prefix/description/

/*
考察：字符串处理、循环嵌套

题目：查找字符串数组中的最长公共前缀
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var prefix string = strs[0]
	if prefix == "" {
		return ""
	}
	for i := 1; i < len(strs); i++ {
		fmt.Println("当前比较的字符串----：", strs[i])

		if strs[i] == "" {
			return ""
		}
		for j := 0; j < len(prefix) && j < len(strs[i]); j++ {
			fmt.Println("当前比较的字符：", prefix[j], strs[i][j], string(prefix[j]), string(strs[i][j]), len(prefix), len(strs[i]))

			if prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				fmt.Println("当前的前缀是：", prefix)
				break
			}
		}
		if len(prefix) > len(strs[i]) {
			prefix = strs[i]
		}
	}
	return prefix
}

func main() {
	// fmt.Println(longestCommonPrefix([]string{"flo是wer", "flo是ight", "floæw"}))
	// fmt.Println(longestCommonPrefix([]string{"flo是w", "flo是wer", "flo是ight"}))
	// fmt.Println(longestCommonPrefix([]string{"flo是w", "flo是werr"}))
	// fmt.Println(longestCommonPrefix([]string{}))
	// fmt.Println(longestCommonPrefix([]string{"fflo", "cflo"}))
	// fmt.Println(longestCommonPrefix([]string{"", ""}))
	// fmt.Println(longestCommonPrefix([]string{"s", ""}))
	// fmt.Println(longestCommonPrefix([]string{"", "s"}))
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
}
