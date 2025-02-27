package models

type Status struct {
	ID   int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Name string `gorm:"type:varchar(255);uniqueIndex;not null" json:"name"`
}

func (Status) TableName() string {
	return "statuses"
}
