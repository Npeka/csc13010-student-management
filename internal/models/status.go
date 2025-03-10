package models

type Status struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;" json:"id"`
	Name string `gorm:"type:varchar(255);uniqueIndex;not null" json:"name"`
}

func (Status) TableName() string {
	return "statuses"
}

type StatusTransition struct {
	ID              uint   `gorm:"primaryKey;autoIncrement;" json:"id"`
	CurrentStatusID uint   `gorm:"not null" json:"current_status_id"`
	CurrentStatus   Status `gorm:"foreignKey:CurrentStatusID;references:ID" json:"current_status"`
	NewStatusID     uint   `gorm:"not null" json:"new_status_id"`
	NewStatus       Status `gorm:"foreignKey:NewStatusID;references:ID" json:"new_status"`
}

func (StatusTransition) TableName() string {
	return "status_transitions"
}
