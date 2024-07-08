package models

// Pod 구조체는 클러스터의 팟을 나타냅니다.
type Pod struct {
    ID     string `bson:"_id,omitempty" json:"id"`
    NodeID string `bson:"node_id" json:"node_id"`
    Status string `bson:"status" json:"status"`
}
