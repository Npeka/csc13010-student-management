package models

type Status struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;" json:"id"`
	Name string `gorm:"type:varchar(255);uniqueIndex;not null" json:"name"`
}

func (Status) TableName() string {
	return "statuses"
}

type StatusTransition struct {
	ID              uint   `gorm:"primaryKey"`
	CurrentStatusID uint   `gorm:"not null"`
	NewStatusID     uint   `gorm:"not null"`
	AllowedRoles    string `gorm:"type:varchar(255);not null"`
}

func (StatusTransition) TableName() string {
	return "status_transitions"
}
