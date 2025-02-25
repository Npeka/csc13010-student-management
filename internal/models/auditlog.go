package models

import "time"

type Action string

const (
	ActionCreate Action = "CREATE"
	ActionUpdate Action = "UPDATE"
	ActionDelete Action = "DELETE"
)

type AuditLog struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	TableName     string    `gorm:"type:varchar(255);not null" json:"table_name"`
	RecordID      int       `gorm:"not null" json:"record_id"`
	Action        Action    `gorm:"type:varchar(10);not null" json:"action"`
	ChangedFields string    `gorm:"type:text;not null" json:"changed_fields"`
	ChangedBy     Role      `gorm:"type:varchar(255);not null" json:"changed_by"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
}
