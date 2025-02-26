package student

import "github.com/gin-gonic/gin"

type IStudentHandlers interface {
	GetStudents() gin.HandlerFunc
	GetStudentByStudentID() gin.HandlerFunc
	CreateStudent() gin.HandlerFunc
	UpdateStudent() gin.HandlerFunc
	DeleteStudent() gin.HandlerFunc
	ImportStudents() gin.HandlerFunc
	ExportStudents() gin.HandlerFunc

	GetOptions() gin.HandlerFunc

	GetFaculties() gin.HandlerFunc
	GetPrograms() gin.HandlerFunc
	GetStatuses() gin.HandlerFunc

	CreateFaculty() gin.HandlerFunc
	CreateProgram() gin.HandlerFunc
	CreateStatus() gin.HandlerFunc

	DeleteFaculty() gin.HandlerFunc
	DeleteProgram() gin.HandlerFunc
	DeleteStatus() gin.HandlerFunc
}
