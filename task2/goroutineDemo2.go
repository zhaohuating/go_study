package main

import (
	"fmt"
	"time"
)

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。

// 任务
type Task struct {
	ID        int
	task      func()
	startTime time.Time
	endTime   time.Time
}

// 任务调度器
type TaskScheduler struct {
	TaskSli []*Task
}

// 获取对象
func NewIntervals() *TaskScheduler {
	return &TaskScheduler{
		TaskSli: make([]*Task, 0),
	}
}

// 添加任务
func (ts *TaskScheduler) add(task *Task) {
	ts.TaskSli = append(ts.TaskSli, task)
}

// 执行任务
func (ts *TaskScheduler) run() {
	stateChan := make(chan bool, 10)
	index := 0
	for i, task := range ts.TaskSli {
		if task != nil {
			go func(t *Task) {
				defer func() {
					stateChan <- true
				}()
				t.startTime = time.Now()
				t.task()
				t.endTime = time.Now()
			}(task)
			index = i
		}
	}

	for {
		if index == len(stateChan)-1 {
			close(stateChan)
			break
		}
	}
}

func main() {
	formatStr := "2006-01-02 15:04:05"
	taskScheduler := NewIntervals()
	taskScheduler.add(getTestTask(1, 2*time.Second)) // 添加任务
	taskScheduler.add(getTestTask(2, 1*time.Second))
	taskScheduler.run()

	for _, t := range taskScheduler.TaskSli {
		fmt.Printf("任务id：%d\n", t.ID)
		fmt.Printf("开始时间：%s\n", t.startTime.Format(formatStr))
		fmt.Printf("结束时间：%s\n", t.endTime.Format(formatStr))
		fmt.Printf("执行耗时：%.f s\n", t.endTime.Sub(t.startTime).Seconds())
	}
}

// 获取模拟任务
func getTestTask(id int, duration time.Duration) *Task {
	return &Task{
		ID: id,
		task: func() {
			time.Sleep(duration)
		},
	}
}
