package status

import "github.com/gin-gonic/gin"

type IStatusHandlers interface {
	GetStatuses() gin.HandlerFunc
	CreateStatus() gin.HandlerFunc
	UpdateStatus() gin.HandlerFunc
	DeleteStatus() gin.HandlerFunc
}
