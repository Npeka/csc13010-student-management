package http

import (
	"github.com/csc13010-student-management/internal/faculty"
	"github.com/gin-gonic/gin"
)

func MapFacultyHandlers(ftGroup *gin.RouterGroup, h faculty.IFacultyHandlers) {
	ftGroup.GET("/", h.GetFaculties())
	ftGroup.POST("/", h.CreateFaculty())
	ftGroup.PUT("/:id", h.UpdateFaculty())
	ftGroup.DELETE("/:id", h.DeleteFaculty())
}
