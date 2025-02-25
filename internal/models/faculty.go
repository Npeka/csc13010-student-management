package models

type Faculty struct {
	Id   int    `gorm:"primaryKey;autoIncrement;not null"`
	Name string `gorm:"type:varchar(255);not null"`
}

func (Faculty) TableName() string {
	return "faculties"
}
