package http

import (
	"github.com/csc13010-student-management/internal/student"
	"github.com/gin-gonic/gin"
)

func MapStudentHandlers(stGroup *gin.RouterGroup, h student.IStudentHandlers) {
	stGroup.GET("", h.GetStudents())
	stGroup.GET("/:student_id", h.GetStudentByStudentID())
	stGroup.POST("", h.CreateStudent())
	stGroup.PATCH("/:student_id", h.UpdateStudent())
	stGroup.DELETE("/:student_id", h.DeleteStudent())
	stGroup.GET("/options", h.GetOptions())
	stGroup.POST("/import", h.ImportStudents())
	// /api/students/export?format=csv
	stGroup.GET("/export", h.ExportStudents())
}
