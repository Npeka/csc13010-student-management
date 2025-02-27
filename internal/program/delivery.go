package program

import "github.com/gin-gonic/gin"

type IProgramHandlers interface {
	GetPrograms() gin.HandlerFunc
	CreateProgram() gin.HandlerFunc
	DeleteProgram() gin.HandlerFunc
}
