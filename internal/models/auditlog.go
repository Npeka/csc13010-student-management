package models

import (
	"time"
)

type Action string

const (
	ActionCreate Action = "CREATE"
	ActionUpdate Action = "UPDATE"
	ActionDelete Action = "DELETE"
)

type AuditLog struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	TableName    string    `gorm:"type:varchar(255);not null" json:"table_name"`
	RecordID     uint      `gorm:"not null" json:"record_id"`
	Action       Action    `gorm:"type:varchar(10);not null" json:"action"`
	FieldChanges string    `gorm:"type:text" json:"field_changes,omitempty"`
	OldRecord    string    `gorm:"type:jsonb" json:"old_record,omitempty"`
	NewRecord    string    `gorm:"type:jsonb" json:"new_record,omitempty"`
	Transaction  string    `gorm:"type:varchar(50)" json:"transaction,omitempty"`
	LSN          int       `gorm:"not null" json:"lsn"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}
