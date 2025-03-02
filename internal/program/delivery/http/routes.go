package http

import (
	"github.com/csc13010-student-management/internal/program"
	"github.com/gin-gonic/gin"
)

func MapProgramHandlers(pgGroup *gin.RouterGroup, ph program.IProgramHandlers) {
	pgGroup.GET("/", ph.GetPrograms())
	pgGroup.POST("/", ph.CreateProgram())
	pgGroup.PUT("/:id", ph.UpdateProgram())
	pgGroup.DELETE("/:id", ph.DeleteProgram())
}
