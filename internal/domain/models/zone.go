package models

type Zone struct {
	Id       int     `json:"zone_id" db:"id"`
	Workload float32 `json:"workload" db:"workload"` // 0 - 1
}
