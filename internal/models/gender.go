package models

type Gender struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (Gender) TableName() string {
	return "genders"
}
