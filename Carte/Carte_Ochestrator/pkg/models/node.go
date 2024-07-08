package models

// Node 구조체는 클러스터의 노드를 나타냅니다.
type Node struct {
    ID       string `bson:"_id,omitempty" json:"id"`
    Hostname string `bson:"hostname" json:"hostname"`
    Status   string `bson:"status" json:"status"`
}
