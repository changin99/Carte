package worker

import "log"

// Task 구조체는 작업 정보를 나타냅니다.
type Task struct {
    ID    string `json:"id"`
    Data  string `json:"data"`
}

// TaskQueue 구조체는 작업 큐를 나타냅니다.
type TaskQueue struct {
    queue chan Task
}

// NewTaskQueue는 새로운 작업 큐를 생성합니다.
func NewTaskQueue(size int) *TaskQueue {
    return &TaskQueue{
        queue: make(chan Task, size),
    }
}

// AddTask는 작업 큐에 Task를 추가합니다.
func (tq *TaskQueue) AddTask(task Task) {
    tq.queue <- task
}

// Run은 작업 큐에서 Task를 처리합니다.
func (tq *TaskQueue) Run() {
    for task := range tq.queue {
        log.Printf("Processing task: %s", task.ID)
        // 여기에 실제 작업 처리 로직을 추가
    }
}
