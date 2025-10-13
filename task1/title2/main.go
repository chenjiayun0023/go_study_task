package main

import "fmt"

//https://leetcode.cn/problems/palindrome-number/description/

/*
9. 回文数：判断一个整数是否是回文数

给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如，121 是回文，而 123 不是。
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	//倒序读取的值
	var arr []int
	for x != 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	fmt.Println("arr的值：", arr)
	for i := 0; i < len(arr)/2; i++ {
		//对称的两边数字是否相等
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	//保存原始值
	var originalX int = x
	//倒序读取的值
	var arr []int
	for x != 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	fmt.Println("arr的值：", arr)

	//倒序后的值
	newX := 0
	for _, v := range arr {
		newX = newX*10 + v // 逐位拼接
		// fmt.Println(newX)
	}
	fmt.Println("newX的值：", newX)

	fmt.Printf("正序的值：%d, 倒序的值： %d \n", originalX, newX)
	return newX == originalX
}

func main() {
	var x int = 0 //12345  12321   123321 0  -2  20  201 202
	fmt.Println(isPalindrome2(x))
}
