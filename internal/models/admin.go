package models

type Admin struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	UserID   uint   `gorm:"uniqueIndex;not null" json:"user_id"`
	User     User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	StaffID  string `gorm:"type:varchar(12);uniqueIndex;not null" json:"staff_id"`
	FullName string `gorm:"type:varchar(100);not null" json:"full_name"`
}
