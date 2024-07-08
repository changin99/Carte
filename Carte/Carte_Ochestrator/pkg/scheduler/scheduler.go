package scheduler

import (
    "Carte_Orchestrator/pkg/models"
    "log"
)

// Scheduler 구조체는 스케줄러를 나타냅니다.
type Scheduler struct {
    Nodes []models.Node
}

// NewScheduler는 새로운 스케줄러를 생성합니다.
func NewScheduler(nodes []models.Node) *Scheduler {
    return &Scheduler{
        Nodes: nodes,
    }
}

// SchedulePod는 팟을 노드에 스케줄링합니다.
func (s *Scheduler) SchedulePod(pod *models.Pod) {
    // 간단한 라운드 로빈 스케줄링 예제
    for _, node := range s.Nodes {
        if node.Status == "Ready" {
            pod.NodeID = node.ID
            log.Printf("Pod %s scheduled to Node %s", pod.ID, node.ID)
            return
        }
    }
    log.Printf("No available nodes to schedule Pod %s", pod.ID)
}
