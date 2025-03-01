package faculty

import "github.com/gin-gonic/gin"

type IFacultyHandlers interface {
	GetFaculties() gin.HandlerFunc
	CreateFaculty() gin.HandlerFunc
	UpdateFaculty() gin.HandlerFunc
	DeleteFaculty() gin.HandlerFunc
}
