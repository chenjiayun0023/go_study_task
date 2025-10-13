package main

import "fmt"

//https://leetcode.cn/problems/valid-parentheses/description/

/*
考察：字符串处理、栈的使用

题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
*/

func isValid(str string) bool {
	//栈
	var stack []rune
	//映射关系
	var mapping = map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
		'否': '是',
	}
	for _, value := range str {
		// fmt.Println("str[i]的值：", value, rune(value))
		if val, ok := mapping[rune(value)]; ok {
			// 如果是右括号，检查栈顶元素是否匹配
			if len(stack) == 0 || stack[len(stack)-1] != val {
				return false
			}
			stack = stack[:len(stack)-1] // 弹出栈顶元素
		} else {
			// 如果是左括号，入栈
			stack = append(stack, rune(value))
		}
	}
	return str != "" && len(stack) == 0
}

func main() {
	var s string = "((())){}" //()[]{}true   ([]){}true   }{false  否是false  是否true   ([)]false  ((()))true   空字符串 false
	fmt.Println(isValid(s))

}
