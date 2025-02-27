package http

import (
	"github.com/csc13010-student-management/internal/faculty"
	"github.com/gin-gonic/gin"
)

func MapfacultyHandlers(ftGroup *gin.RouterGroup, h faculty.IFacultyHandlers) {
	ftGroup.GET("/", h.GetFaculties())
	ftGroup.POST("/", h.CreateFaculty())
	ftGroup.DELETE("/:id", h.DeleteFaculty())
}
