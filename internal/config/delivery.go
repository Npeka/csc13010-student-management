package config

import "github.com/gin-gonic/gin"

type IConfigHandlers interface {
	GetConfig() gin.HandlerFunc
	UpdateConfig() gin.HandlerFunc
}
