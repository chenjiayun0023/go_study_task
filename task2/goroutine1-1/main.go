package main

//Goroutine

/*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/

import (
	"fmt"
	"sync"
)

func oddPrint() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("奇数:", i)
	}
}

func evenPrint() {
	for i := 2; i <= 10; i += 2 {
		fmt.Println("偶数:", i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		oddPrint()
	}()

	go func() {
		defer wg.Done()
		evenPrint()
	}()

	wg.Wait()
	fmt.Println("main函数结束")
}
