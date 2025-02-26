package models

type Faculty struct {
	Id   int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Name string `gorm:"type:varchar(255);uniqueIndex;not null" json:"name"`
}

func (Faculty) TableName() string {
	return "faculties"
}
