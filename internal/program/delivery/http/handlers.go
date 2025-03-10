package http

import (
	"net/http"
	"strconv"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/program"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/csc13010-student-management/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

type programHandlers struct {
	pu program.IProgramUsecase
	lg *logger.LoggerZap
}

func NewProgramHandlers(
	pu program.IProgramUsecase,
	lg *logger.LoggerZap,
) program.IProgramHandlers {
	return &programHandlers{
		pu: pu,
		lg: lg,
	}
}

func (ph *programHandlers) GetPrograms() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "program.GetPrograms")
		defer span.Finish()

		programs, err := ph.pu.GetPrograms(ctx)
		if err != nil {
			logger.LogResponseError(c, ph.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, http.StatusOK, programs)
	}
}

func (ph *programHandlers) CreateProgram() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "program.CreateProgram")
		defer span.Finish()

		var program models.Program
		if err := c.ShouldBindJSON(&program); err != nil {
			logger.ErrResponseWithLog(c, ph.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		err := ph.pu.CreateProgram(ctx, &program)
		if err != nil {
			logger.LogResponseError(c, ph.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, http.StatusCreated, program)
	}
}

func (ph *programHandlers) UpdateProgram() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "program.UpdateProgram")
		defer span.Finish()

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			logger.ErrResponseWithLog(c, ph.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		var program models.Program
		if err := c.ShouldBindJSON(&program); err != nil {
			logger.LogResponseError(c, ph.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		program.ID = uint(id)
		err = ph.pu.UpdateProgram(ctx, &program)
		if err != nil {
			logger.LogResponseError(c, ph.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, http.StatusOK, program)
	}
}

func (ph *programHandlers) DeleteProgram() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "program.DeleteProgram")
		defer span.Finish()

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			logger.ErrResponseWithLog(c, ph.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		err = ph.pu.DeleteProgram(ctx, uint(id))
		if err != nil {
			logger.ErrResponseWithLog(c, ph.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, http.StatusNoContent, nil)
	}
}
