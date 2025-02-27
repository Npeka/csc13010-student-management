package student

import "github.com/gin-gonic/gin"

type IStudentHandlers interface {
	GetStudents() gin.HandlerFunc
	GetStudentByStudentID() gin.HandlerFunc
	GetFullInfoStudentByStudentID() gin.HandlerFunc
	CreateStudent() gin.HandlerFunc
	UpdateStudent() gin.HandlerFunc
	DeleteStudent() gin.HandlerFunc
	GetOptions() gin.HandlerFunc
}
