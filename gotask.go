// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
//   - 考察点 ：协程原理、并发任务调度。
package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 定义任务类型
type Task func()

// TaskResult 存储任务执行结果
type TaskResult struct {
	ID       int
	Duration time.Duration
}

type Scheduler struct {
	tasks      []Task
	concurrent int
}

// NewScheduler 创建调度器实例
func NewScheduler(concurrent int) *Scheduler {
	return &Scheduler{
		concurrent: concurrent,
	}
}

// AddTask 添加任务
func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

// Run 并发执行所有任务并返回统计结果
func (s *Scheduler) Run() map[int]time.Duration {
	var wg sync.WaitGroup
	resultChan := make(chan TaskResult, len(s.tasks))
	results := make(map[int]time.Duration)

	//创建令牌桶控制并发数
	sem := make(chan struct{}, s.concurrent)

	for i, task := range s.tasks {
		wg.Add(1)
		go func(id int, t Task) {
			defer wg.Done()
			sem <- struct{}{} //获取令牌

			start := time.Now()
			t() //执行任务
			duration := time.Since(start)

			resultChan <- TaskResult{ID: id, Duration: duration}
			<-sem //释放令牌
		}(i, task)
	}

	//收集结果
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for res := range resultChan {
		results[res.ID] = res.Duration
	}

	return results
}

func main() {
	scheduler := NewScheduler(3) //最大并发数3

	//添加示例任务
	scheduler.AddTask(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Task 1 completed")
	})

	scheduler.AddTask(func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Task 2 completed")
	})

	scheduler.AddTask(func() {
		time.Sleep(500 * time.Millisecond) //500毫秒
		fmt.Println("Task 3 completed")
	})

	scheduler.AddTask(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Task 4 completed")
	})

	//执行并获取结果
	results := scheduler.Run()

	//打印结果
	fmt.Println("\nExecution Statistics:")
	for id, duration := range results {
		fmt.Printf("Task %d took %v\n", id+1, duration)
	}
}
