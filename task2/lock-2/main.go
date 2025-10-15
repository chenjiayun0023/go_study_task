package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//锁机制

/*
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/

// 无锁计数器结构
type LockFreeCounter struct {
	value int64
}

// 创建新计数器
func NewCounter() *LockFreeCounter {
	return &LockFreeCounter{}
}

// 增加计数器（并发安全）
func (c *LockFreeCounter) Add(delta int64) {
	atomic.AddInt64(&c.value, delta)
}

// 获取当前值（并发安全）
func (c *LockFreeCounter) Get() int64 {
	return atomic.LoadInt64(&c.value)
}

// 重置计数器
func (c *LockFreeCounter) Reset() {
	atomic.StoreInt64(&c.value, 0)
}

func main() {
	counter := NewCounter()

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Add(1)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("最终值: %d\n", counter.Get())

	counter.Reset()
	fmt.Printf("重置后: %d\n", counter.Get())
}
