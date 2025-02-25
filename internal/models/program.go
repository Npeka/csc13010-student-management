package models

type Program struct {
	Id   int    `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}

func (Program) TableName() string {
	return "programs"
}
