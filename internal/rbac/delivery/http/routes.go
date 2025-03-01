package http

import (
	"github.com/csc13010-student-management/internal/rbac"
	"github.com/gin-gonic/gin"
)

func MaprbacHandlers(rlGroup *gin.RouterGroup, h rbac.IRbacHandlers) {
	rlGroup.POST("/role", h.AddRole())
	rlGroup.DELETE("/role", h.DeleteRole())

	rlGroup.POST("/role/api", h.AddAPIForRole())
	rlGroup.DELETE("/role/api", h.DeleteAPIForRole())

	rlGroup.POST("/role/user", h.AddRoleForUser())
	rlGroup.DELETE("/role/user", h.DeleteRoleForUser())

	rlGroup.POST("/role/api/role", h.AddRoleForAPI())
	rlGroup.DELETE("/role/api/role", h.DeleteRoleForAPI())

	rlGroup.GET("/auth", h.CheckAuth())
	rlGroup.GET("/notification", h.Notification())
}
