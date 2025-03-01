package models

type Gender struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;" json:"id"`
	Name string `gorm:"type:varchar(255);uniqueIndex;not null" json:"name"`
}

func (Gender) TableName() string {
	return "genders"
}
