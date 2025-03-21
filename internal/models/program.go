package models

type Program struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;" json:"id"`
	Name string `gorm:"type:varchar(255);uniqueIndex;not null" json:"name"`
}

func (Program) TableName() string {
	return "programs"
}
