package main

import "fmt"

//指针

/*
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
*/
func add(numPtr *int) {
	*numPtr += 10
}

/*
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/

func multiply(numPtr *[]int) {
	for i := range *numPtr {
		(*numPtr)[i] *= 2
	}
}

func main() {
	num := 20
	add(&num)
	fmt.Println(num) // 输出 30

	num1 := []int{1, 2, 3, 4, 5}
	numPtr := &num1
	multiply(numPtr)
	fmt.Println(num1) // 输出 2 4
}
