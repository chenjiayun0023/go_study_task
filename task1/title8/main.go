package main

import "fmt"

//https://leetcode.cn/problems/two-sum/

/*
考察：数组遍历、map使用

题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数的下标
*/
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		fmt.Println("i、num、numMap: ", i, num, numMap)

		if j, ok := numMap[target-num]; ok {
			fmt.Println("j、ok: ", j, ok)
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result)
}
