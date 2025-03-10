package models

type Config struct {
	ID            uint `gorm:"primaryKey,autoIncrement" json:"id"`
	EmailDomain   bool `gorm:"not null" json:"email_domain"`
	ValidatePhone bool `gorm:"not null" json:"validate_phone"`
	StatusRules   bool `gorm:"not null" json:"status_rules"`
	DeleteLimit   bool `gorm:"not null" json:"delete_limit"`
}

func (Config) TableName() string {
	return "configs"
}
