package main

import (
	"fmt"
	"sync"
)

//锁机制

/*
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，
每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) add() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count += 1
}

func (c *Counter) getCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count

}

func main() {
	counter := &Counter{}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.add()
			}
		}()
	}

	wg.Wait()
	fmt.Println("计数：", counter.getCount())
}
