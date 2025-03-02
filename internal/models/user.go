package models

import (
	"time"
)

const (
	RoleUser    = "user"
	RoleAdmin   = "admin"
	RoleStudent = "student"
	RoleTeacher = "teacher"
)

type Role struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(50);unique;not null" json:"name"`
}

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Username  string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"username"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	RoleId    uint      `gorm:"not null" json:"role_id"`
	Role      Role      `gorm:"foreignKey:RoleId" json:"role,omitempty"`
	CreatedAt time.Time `gorm:"type:timestamp; default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp; default:CURRENT_TIMESTAMP"`
}

func (User) TableName() string {
	return "users"
}

// // AfterCreate hooks into GORM to link user with Casbin role after creation
// func (u *User) AfterCreate(tx *gorm.DB) error {
// 	// Load role name from database for this user
// 	var role Role
// 	if err := tx.Model(&Role{}).Where("id = ?", u.RoleID).First(&role).Error; err != nil {
// 		return fmt.Errorf("failed to get role for user %d: %w", u.ID, err)
// 	}

// 	// Get Casbin enforcer from context or create one
// 	enforcer, err := getCasbinEnforcer()
// 	if err != nil {
// 		return fmt.Errorf("failed to get casbin enforcer: %w", err)
// 	}

// 	// Add role for user in Casbin
// 	// Format: g, user_identifier, role_name
// 	_, err = enforcer.AddGroupingPolicy(fmt.Sprintf("user:%d", u.ID), role.Name)
// 	if err != nil {
// 		return fmt.Errorf("failed to add user role in casbin: %w", err)
// 	}

// 	return nil
// }

// // AfterUpdate hooks into GORM to update Casbin role if user's role has changed
// func (u *User) AfterUpdate(tx *gorm.DB) error {
// 	// Check if RoleID was changed
// 	var oldUser User
// 	if err := tx.Unscoped().Where("id = ?", u.ID).First(&oldUser).Error; err != nil {
// 		return err
// 	}

// 	// If role wasn't changed, no need to update Casbin
// 	if oldUser.RoleID == u.RoleID {
// 		return nil
// 	}

// 	// Load old and new role names
// 	var oldRole, newRole Role
// 	if err := tx.Model(&Role{}).Where("id = ?", oldUser.RoleID).First(&oldRole).Error; err != nil {
// 		return fmt.Errorf("failed to get old role: %w", err)
// 	}
// 	if err := tx.Model(&Role{}).Where("id = ?", u.RoleID).First(&newRole).Error; err != nil {
// 		return fmt.Errorf("failed to get new role: %w", err)
// 	}

// 	// Get Casbin enforcer
// 	enforcer, err := getCasbinEnforcer()
// 	if err != nil {
// 		return fmt.Errorf("failed to get casbin enforcer: %w", err)
// 	}

// 	// Remove old role assignment
// 	_, err = enforcer.RemoveGroupingPolicy(fmt.Sprintf("user:%d", u.ID), oldRole.Name)
// 	if err != nil {
// 		return fmt.Errorf("failed to remove old role in casbin: %w", err)
// 	}

// 	// Add new role assignment
// 	_, err = enforcer.AddGroupingPolicy(fmt.Sprintf("user:%d", u.ID), newRole.Name)
// 	if err != nil {
// 		return fmt.Errorf("failed to add new role in casbin: %w", err)
// 	}

// 	return nil
// }

// // AfterDelete hooks into GORM to remove user from Casbin when user is deleted
// func (u *User) AfterDelete(tx *gorm.DB) error {
// 	// Get Casbin enforcer
// 	enforcer, err := getCasbinEnforcer()
// 	if err != nil {
// 		return fmt.Errorf("failed to get casbin enforcer: %w", err)
// 	}

// 	// Remove all role assignments for this user
// 	_, err = enforcer.RemoveFilteredGroupingPolicy(0, fmt.Sprintf("user:%d", u.ID))
// 	if err != nil {
// 		return fmt.Errorf("failed to remove user from casbin: %w", err)
// 	}

// 	return nil
// }

// // Helper function to get or initialize Casbin enforcer
// // You'll need to adapt this based on how you're managing the enforcer in your application
// func getCasbinEnforcer() (*casbin.Enforcer, error) {
// 	// This is a placeholder - in a real application, you might:
// 	// 1. Get the enforcer from a global variable
// 	// 2. Get it from a service locator/dependency injection container
// 	// 3. Create a new one each time (not recommended for performance reasons)

// 	// Example implementation (replace with your actual approach):
// 	return casbin.NewEnforcer("path/to/model.conf", "path/to/policy.csv")
// }

// // Function to check if a user has permission to perform an action on a resource
// func CheckUserPermission(userID uint, resource string, action string) (bool, error) {
// 	enforcer, err := getCasbinEnforcer()
// 	if err != nil {
// 		return false, err
// 	}

// 	// Format: sub, obj, act
// 	return enforcer.Enforce(fmt.Sprintf("user:%d", userID), resource, action)
// }

// // Function to sync all users with Casbin (useful after migration or system reset)
// func SyncAllUsersWithCasbin(db *gorm.DB) error {
// 	enforcer, err := getCasbinEnforcer()
// 	if err != nil {
// 		return err
// 	}

// 	// Clear all user role assignments in Casbin
// 	_, err = enforcer.RemoveFilteredGroupingPolicy(0, "user:")
// 	if err != nil {
// 		return err
// 	}

// 	// Get all users with their roles
// 	var users []struct {
// 		UserID   uint
// 		RoleName string
// 	}

// 	err = db.Table("users").
// 		Select("users.id as user_id, roles.name as role_name").
// 		Joins("JOIN roles ON users.role_id = roles.id").
// 		Find(&users).Error

// 	if err != nil {
// 		return err
// 	}

// 	// Add all users to Casbin
// 	for _, user := range users {
// 		_, err = enforcer.AddGroupingPolicy(fmt.Sprintf("user:%d", user.UserID), user.RoleName)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
