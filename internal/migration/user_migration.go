package migration

import (
	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/pkg/utils/crypto"
	"gorm.io/gorm"
)

func seedUsers(db *gorm.DB) {
	users := []models.User{
		{
			Username: "22127180",
			Password: crypto.GetHash("22127180"),
			RoleId:   2,
		},
		{
			Username: "22127108",
			Password: crypto.GetHash("22127108"),
			RoleId:   2,
		},
		{
			Username: "22127419",
			Password: crypto.GetHash("22127419"),
			RoleId:   2,
		},
	}

	db.Table("users").Create(&users)
}
