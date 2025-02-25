package student

import "github.com/gin-gonic/gin"

type IStudentHandlers interface {
	GetStudents() gin.HandlerFunc
	CreateStudent() gin.HandlerFunc
	UpdateStudent() gin.HandlerFunc
	DeleteStudent() gin.HandlerFunc
	GetOptions() gin.HandlerFunc
}
