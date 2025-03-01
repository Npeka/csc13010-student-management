package rbac

import "github.com/gin-gonic/gin"

type IRbacHandlers interface {
	AddRole() gin.HandlerFunc
	AddRoleForUser() gin.HandlerFunc
	DeleteRoleForUser() gin.HandlerFunc
	AddAPIForRole() gin.HandlerFunc
	DeleteAPIForRole() gin.HandlerFunc
	AddRoleForAPI() gin.HandlerFunc
	DeleteRoleForAPI() gin.HandlerFunc
	DeleteRole() gin.HandlerFunc
	CheckAuth() gin.HandlerFunc
	Notification() gin.HandlerFunc
}
