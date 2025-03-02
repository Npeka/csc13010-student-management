package http

import (
	"github.com/csc13010-student-management/internal/middleware"
	"github.com/csc13010-student-management/internal/student"
	"github.com/gin-gonic/gin"
)

// MapStudentHandlers ánh xạ API student vào router
func MapStudentHandlers(stGroup *gin.RouterGroup, h student.IStudentHandlers, mw *middleware.MiddlewareManager) {
	// Các API cần quyền "read"
	// stGroup.GET("/", mw.RBAC("/students", "read"), h.GetStudents())
	// stGroup.GET("/:student_id", mw.RBAC("/students/:student_id", "read"), h.GetStudentByStudentID())
	// stGroup.GET("/full/:student_id", mw.RBAC("/students/:student_id", "read"), h.GetFullInfoStudentByStudentID())

	// // Các API cần quyền "write"
	// stGroup.POST("/", mw.RBAC("/students", "write"), h.CreateStudent())
	// stGroup.PATCH("/:student_id", mw.RBAC("/students/:student_id", "write"), h.UpdateStudent())
	// stGroup.DELETE("/:student_id", mw.RBAC("/students/:student_id", "write"), h.DeleteStudent())

	// // API không cần RBAC
	// stGroup.GET("/options", h.GetOptions())

	// route not apply middleware
	stGroup.GET("/", h.GetStudents())
	stGroup.GET("/:student_id", h.GetStudentByStudentID())
	stGroup.GET("/full/:student_id", h.GetFullInfoStudentByStudentID())

	// // Các API cần quyền "write"
	stGroup.POST("/", h.CreateStudent())
	stGroup.PATCH("/:student_id", h.UpdateStudent())
	stGroup.DELETE("/:student_id", h.DeleteStudent())

	// // API không cần RBAC
	stGroup.GET("/options", h.GetOptions())
}
