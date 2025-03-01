package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Action string

const (
	ActionCreate Action = "CREATE"
	ActionUpdate Action = "UPDATE"
	ActionDelete Action = "DELETE"
)

type FieldChange struct {
	Field    string      `json:"field"`
	OldValue interface{} `json:"old_value,omitempty"`
	NewValue interface{} `json:"new_value"`
}
type AuditLog struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	TableName   string        `gorm:"type:varchar(255);not null" json:"table_name"`
	RecordID    uint          `gorm:"not null" json:"record_id"`
	Action      Action        `gorm:"type:varchar(10);not null" json:"action"`
	Changes     []FieldChange `gorm:"-" json:"changes"`
	ChangesJSON string        `gorm:"type:text;not null;column:changes" json:"-"`
	UserID      uint          `gorm:"not null" json:"user_id"`
	UserEmail   string        `gorm:"type:varchar(255);not null" json:"user_email"`
	UserRole    string        `gorm:"type:varchar(50);not null" json:"user_role"`
	CreatedAt   time.Time     `gorm:"autoCreateTime" json:"created_at"`
	IPAddress   string        `gorm:"type:varchar(45)" json:"ip_address,omitempty"`
	UserAgent   string        `gorm:"type:text" json:"user_agent,omitempty"`
}

func (a *AuditLog) BeforeSave(tx *gorm.DB) error {
	jsonData, err := json.Marshal(a.Changes)
	if err != nil {
		return err
	}
	a.ChangesJSON = string(jsonData)
	return nil
}

func (a *AuditLog) AfterFind(tx *gorm.DB) error {
	if a.ChangesJSON != "" {
		return json.Unmarshal([]byte(a.ChangesJSON), &a.Changes)
	}
	return nil
}

func LogModelChanges(
	tx *gorm.DB,
	tableName string,
	recordID uint,
	action Action,
	oldModel, newModel interface{},
	userID uint,
	userEmail,
	userRole string,
) error {
	var changes []FieldChange

	if action == ActionCreate {
		// For creation, log all non-zero fields
		newValues := make(map[string]interface{})
		newModelBytes, _ := json.Marshal(newModel)
		json.Unmarshal(newModelBytes, &newValues)

		for field, value := range newValues {
			if field != "ID" && field != "CreatedAt" && field != "UpdatedAt" {
				changes = append(changes, FieldChange{
					Field:    field,
					NewValue: value,
				})
			}
		}
	} else if action == ActionUpdate {
		// For updates, compare old and new values
		oldValues := make(map[string]interface{})
		newValues := make(map[string]interface{})
		oldModelBytes, _ := json.Marshal(oldModel)
		newModelBytes, _ := json.Marshal(newModel)
		json.Unmarshal(oldModelBytes, &oldValues)
		json.Unmarshal(newModelBytes, &newValues)

		for field, newValue := range newValues {
			if field != "ID" && field != "CreatedAt" && field != "UpdatedAt" {
				oldValue, exists := oldValues[field]
				if !exists || oldValue != newValue {
					changes = append(changes, FieldChange{
						Field:    field,
						OldValue: oldValue,
						NewValue: newValue,
					})
				}
			}
		}
	} else if action == ActionDelete {
		// For deletion, log the entire deleted record
		oldValues := make(map[string]interface{})
		oldModelBytes, _ := json.Marshal(oldModel)
		json.Unmarshal(oldModelBytes, &oldValues)

		for field, value := range oldValues {
			if field != "ID" && field != "CreatedAt" && field != "UpdatedAt" {
				changes = append(changes, FieldChange{
					Field:    field,
					OldValue: value,
				})
			}
		}
	}

	// Skip if no changes
	if len(changes) == 0 && action != ActionDelete {
		return nil
	}

	auditLog := AuditLog{
		TableName: tableName,
		RecordID:  recordID,
		Action:    action,
		Changes:   changes,
		UserID:    userID,
		UserEmail: userEmail,
		UserRole:  userRole,
		CreatedAt: time.Now(),
	}

	return tx.Create(&auditLog).Error
}
