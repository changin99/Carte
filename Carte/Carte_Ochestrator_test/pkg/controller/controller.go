package controller

import (
    "log"
    "Carte_Orchestrator/pkg/models"
    "Carte_Orchestrator/pkg/scheduler"
)

// Controller 구조체는 컨트롤러를 나타냅니다.
type Controller struct {
    Scheduler *scheduler.Scheduler
}

// NewController는 새로운 컨트롤러를 생성합니다.
func NewController(scheduler *scheduler.Scheduler) *Controller {
    return &Controller{
        Scheduler: scheduler,
    }
}

// HandlePodCreation은 팟 생성 요청을 처리합니다.
func (c *Controller) HandlePodCreation(pod *models.Pod) {
    c.Scheduler.SchedulePod(pod)
    log.Printf("Handled pod creation: %s", pod.ID)
}
