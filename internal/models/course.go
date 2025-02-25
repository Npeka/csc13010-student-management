package models

type Course struct {
	Id   int    `gorm:"primaryKey;autoIncrement;not null"`
	Name string `gorm:"type:varchar(255);uniqueIndex;not null"`
}

func (Course) TableName() string {
	return "courses"
}
