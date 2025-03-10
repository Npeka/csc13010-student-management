package http

import (
	"github.com/csc13010-student-management/internal/middleware"
	"github.com/csc13010-student-management/internal/student"
	"github.com/gin-gonic/gin"
)

func MapStudentHandlers(stGroup *gin.RouterGroup, h student.IStudentHandlers, mw *middleware.MiddlewareManager) {
	stGroup.GET("/", h.GetStudents())
	stGroup.GET("/:student_id", h.GetStudentByStudentID())
	stGroup.GET("/full/:student_id", h.GetFullInfoStudentByStudentID())

	stGroup.POST("/", h.CreateStudent())
	stGroup.PUT("/:student_id", h.UpdateStudent())
	stGroup.DELETE("/:student_id", h.DeleteStudent())

	stGroup.GET("/options", h.GetOptions())
	stGroup.GET("/:student_id/export", h.Export())
}
