package api

import (
    "github.com/gin-gonic/gin"
    "Carte_Orchestrator/pkg/worker"
    "net/http"
)

// SetupRouter는 API 서버의 라우터를 설정합니다.
func SetupRouter(taskQueue *worker.TaskQueue) *gin.Engine {
    router := gin.Default()

    // /tasks 엔드포인트 설정
    router.POST("/tasks", func(c *gin.Context) {
        var task worker.Task
        // 요청 바디를 Task 구조체로 바인딩
        if err := c.ShouldBindJSON(&task); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        // 작업 큐에 Task 추가
        taskQueue.AddTask(task)
        c.JSON(http.StatusAccepted, gin.H{"message": "Task accepted"})
    })

    return router
}
