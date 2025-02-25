package student

import "github.com/gin-gonic/gin"

type IStudentHandler interface {
	GetStudents() gin.HandlerFunc
	CreateStudent() gin.HandlerFunc
	UpdateStudent() gin.HandlerFunc
	DeleteStudent() gin.HandlerFunc
	SearchStudents() gin.HandlerFunc
}
