package http

import (
	"github.com/csc13010-student-management/internal/student"
	"github.com/gin-gonic/gin"
)

func MapStudentHandlers(stGroup *gin.RouterGroup, h student.IStudentHandler) {
	stGroup.GET("", h.GetStudents())
	stGroup.POST("", h.CreateStudent())
	stGroup.PATCH("/:student_id", h.UpdateStudent())
	stGroup.DELETE("/:student_id", h.DeleteStudent())
	stGroup.GET("/search", h.SearchStudents())

	stGroup.GET("/options", h.GetOptions())
}
