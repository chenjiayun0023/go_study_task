package main

import (
	"fmt"
	"sync"
	"time"
)

//Goroutine

/*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/

// Task 表示一个可执行的任务类型
type Task func()

// TaskResult 存储任务执行结果
type TaskResult struct {
	ID       int           // 任务ID
	Duration time.Duration // 执行耗时
	Error    error         // 执行错误
}

// Scheduler 任务调度器
type Scheduler struct {
	tasks   []Task         // 任务队列
	results []TaskResult   // 执行结果
	wg      sync.WaitGroup // 等待组
	// mu      sync.Mutex     // 互斥锁保护结果
}

// NewScheduler 创建调度器实例
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make([]Task, 0),
	}
}

// AddTask 添加任务到调度器
func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

// Run 并发执行所有任务并等待完成
func (s *Scheduler) Run() {
	s.results = make([]TaskResult, len(s.tasks))
	s.wg.Add(len(s.tasks))

	for id, task := range s.tasks {
		// 启动协程执行任务
		go s.execute(id, task)
	}

	s.wg.Wait() // 等待所有任务完成
}

// execute 执行单个任务并记录结果
func (s *Scheduler) execute(id int, task Task) {
	defer s.wg.Done()

	start := time.Now()
	defer func() {
		// 捕获可能的panic
		if r := recover(); r != nil {
			// s.mu.Lock()
			s.results[id] = TaskResult{
				ID:       id,
				Duration: time.Since(start),
				Error:    fmt.Errorf("panic: %v", r),
			}
			// s.mu.Unlock()
		}
	}()

	// 执行任务
	task()

	// 记录执行时间
	// s.mu.Lock()
	s.results[id] = TaskResult{
		ID:       id,
		Duration: time.Since(start),
	}
	// s.mu.Unlock()
}

// Results 返回任务执行结果
func (s *Scheduler) Results() []TaskResult {
	return s.results
}

func main() {
	// 创建调度器实例
	scheduler := NewScheduler()

	// 添加示例任务
	scheduler.AddTask(func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Task 1 completed")
	})

	scheduler.AddTask(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Task 2 completed")
	})

	scheduler.AddTask(func() {
		time.Sleep(800 * time.Millisecond)
		fmt.Println("Task 3 completed")
	})

	scheduler.AddTask(func() {
		panic("Simulated panic in task 4")
	})

	// 执行所有任务
	fmt.Println("Starting task execution...")
	startTime := time.Now()
	scheduler.Run()
	fmt.Printf("All tasks completed in %v\n\n", time.Since(startTime))

	// 打印结果
	fmt.Println("Execution results:")
	for _, result := range scheduler.Results() {
		if result.Error != nil {
			fmt.Printf("Task %d: ERROR! %s (took %s)\n",
				result.ID+1, result.Error, result.Duration.Round(time.Millisecond))
		} else {
			fmt.Printf("Task %d: completed in %s\n",
				result.ID+1, result.Duration.Round(time.Millisecond))
		}
	}
}
