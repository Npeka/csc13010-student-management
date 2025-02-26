package models

type Course struct {
	Id   int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Name string `gorm:"type:varchar(255);uniqueIndex;not null" json:"name"`
}

func (Course) TableName() string {
	return "courses"
}
