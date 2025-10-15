package main

import (
	"fmt"
	"sync"
)

//Channel

/*
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/

func send(ch chan<- int) {
	defer close(ch)
	for i := 1; i <= 100; i++ {
		ch <- i
		// fmt.Printf("发送: %d\n", i)
	}
}

func receive(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收: %d\n", v)
	}
}

func main() {
	// 创建一个带缓冲的channel
	ch := make(chan int, 5)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		send(ch)
	}()

	go func() {
		defer wg.Done()
		receive(ch)
	}()

	wg.Wait()
	fmt.Println("接收完成")
}
