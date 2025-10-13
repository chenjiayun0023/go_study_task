package main

import "fmt"

//https://leetcode.cn/problems/single-number/

/*
136. 只出现一次的数字：
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，
然后再遍历 map 找到出现次数为1的元素。
*/

func singleNumber(nums []int) (int, bool) {
	//统计出现的次数
	map1 := make(map[int]int)

	for _, value := range nums {

		//首次出现初始化为1
		if v, ok := map1[value]; !ok {
			fmt.Println("首次出现", v, ok)
			map1[value] = 1
		} else {
			fmt.Println(v, ok)
			map1[value] = map1[value] + 1
		}
	}

	fmt.Println("map1的值：", map1)

	//找出只出现一次的数字
	for key, value := range map1 {
		if value == 1 {
			return key, true //true已找到
		}
	}
	return 0, false //false未找到
}

func main() {
	var nums []int = []int{4, 1, 2, 1, 2}
	// var nums []int = []int{1, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
