package http

import (
	"github.com/csc13010-student-management/internal/auth"
	"github.com/gin-gonic/gin"
)

func MapAuthHandlers(atGroup *gin.RouterGroup, ah auth.IAuthHandlers) {
	atGroup.POST("/login", ah.Login())
	atGroup.POST("/register", ah.Register())
	atGroup.POST("/logout", ah.Logout())
	atGroup.POST("/refresh", ah.Refresh())
}
