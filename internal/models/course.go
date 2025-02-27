package models

type Course struct {
	ID   int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Name string `gorm:"type:varchar(255);uniqueIndex;not null" json:"name"`
}

func (Course) TableName() string {
	return "courses"
}
