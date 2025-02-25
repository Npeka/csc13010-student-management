package models

type Status struct {
	Id   int    `gorm:"primaryKey;autoIncrement;not null"`
	Name string `gorm:"type:varchar(255);uniqueIndex;not null"`
}

func (Status) TableName() string {
	return "statuses"
}
