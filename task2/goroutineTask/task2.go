package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 表示一个任务，包含一个函数和一个名称
type Task struct {
	Name string
	Fn   func()
}

// TaskScheduler 任务调度器
type TaskScheduler struct {
	tasks []Task
}

// NewTaskScheduler 创建一个新的任务调度器
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks: make([]Task, 0),
	}
}

// AddTask 添加一个任务到调度器
func (ts *TaskScheduler) AddTask(name string, fn func()) {
	ts.tasks = append(ts.tasks, Task{Name: name, Fn: fn})
}

// Execute 执行所有任务并统计每个任务的执行时间
func (ts *TaskScheduler) Execute() {
	var wg sync.WaitGroup
	results := make(chan string, len(ts.tasks))

	for _, task := range ts.tasks {
		wg.Add(1)
		go func(task Task) {
			defer wg.Done()
			startTime := time.Now()
			task.Fn()
			duration := time.Since(startTime)
			results <- fmt.Sprintf("任务 %s 执行时间: %v", task.Name, duration)
		}(task)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}

func main() {
	scheduler := NewTaskScheduler()

	// 添加一些示例任务
	scheduler.AddTask("任务1", func() {
		time.Sleep(1 * time.Second)
		fmt.Println("任务1 完成")
	})

	scheduler.AddTask("任务2", func() {
		time.Sleep(2 * time.Second)
		fmt.Println("任务2 完成")
	})

	scheduler.AddTask("任务3", func() {
		time.Sleep(3 * time.Second)
		fmt.Println("任务3 完成")
	})

	// 执行任务调度器
	scheduler.Execute()
}
