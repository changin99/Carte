package main

import (
    "log"
    "Carte_Orchestrator/pkg/api"
    "Carte_Orchestrator/pkg/db"
    "Carte_Orchestrator/pkg/worker"
    "github.com/gin-gonic/gin"
)

func main() {
    // 데이터베이스 초기화
    db.InitMongoDB()

    // 작업 큐 초기화
    taskQueue := worker.NewTaskQueue(10)
    go taskQueue.Run()

    // 라우팅 설정
    router := api.SetupRouter(taskQueue)

    // 서버 시작
    log.Fatal(router.Run(":8080"))
}
