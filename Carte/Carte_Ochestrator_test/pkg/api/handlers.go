package api

import (
    "net/http"
    "Carte_Orchestrator/pkg/worker"
    "github.com/gin-gonic/gin"
)

// TaskHandler는 작업을 처리하는 핸들러입니다.
func TaskHandler(taskQueue *worker.TaskQueue) gin.HandlerFunc {
    return func(c *gin.Context) {
        var task worker.Task
        // 요청 바디를 Task 구조체로 바인딩
        if err := c.ShouldBindJSON(&task); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        // 작업 큐에 Task 추가
        taskQueue.AddTask(task)
        c.JSON(http.StatusAccepted, gin.H{"message": "Task accepted"})
    }
}
