package auth

import "github.com/gin-gonic/gin"

type IAuthHandlers interface {
	Login() gin.HandlerFunc
	Logout() gin.HandlerFunc
	Refresh() gin.HandlerFunc
	Register() gin.HandlerFunc
}
