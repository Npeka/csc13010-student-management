package models

type Faculty struct {
	ID   int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Name string `gorm:"type:varchar(255);uniqueIndex;not null" json:"name"`
}

func (Faculty) TableName() string {
	return "faculties"
}
