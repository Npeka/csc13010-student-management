package http

import (
	"github.com/csc13010-student-management/internal/student"
	"github.com/gin-gonic/gin"
)

func MapStudentHandlers(stGroup *gin.RouterGroup, h student.IStudentHandlers) {
	stGroup.GET("/", h.GetStudents())
	stGroup.GET("/:student_id", h.GetStudentByStudentID())
	stGroup.POST("/", h.CreateStudent())
	stGroup.PATCH("/:student_id", h.UpdateStudent())
	stGroup.DELETE("/:student_id", h.DeleteStudent())
	stGroup.GET("/options", h.GetOptions())
	stGroup.POST("/import", h.ImportStudents())
	stGroup.GET("/export", h.ExportStudents())

	stGroup.GET("/faculties", h.GetFaculties())
	stGroup.GET("/programs", h.GetPrograms())
	stGroup.GET("/statuses", h.GetStatuses())

	stGroup.POST("/faculties", h.CreateFaculty())
	stGroup.POST("/programs", h.CreateProgram())
	stGroup.POST("/statuses", h.CreateStatus())

	stGroup.DELETE("/faculties/:faculty_id", h.DeleteFaculty())
	stGroup.DELETE("/programs/:program_id", h.DeleteProgram())
	stGroup.DELETE("/statuses/:status_id", h.DeleteStatus())
}
