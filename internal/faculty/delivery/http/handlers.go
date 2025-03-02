package http

import (
	"net/http"
	"strconv"

	"github.com/csc13010-student-management/internal/faculty"
	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/csc13010-student-management/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

type facultyHandlers struct {
	fu faculty.IFacultyUsecase
	lg *logger.LoggerZap
}

func NewFacultyHandlers(
	fu faculty.IFacultyUsecase,
	lg *logger.LoggerZap,
) faculty.IFacultyHandlers {
	return &facultyHandlers{
		fu: fu,
		lg: lg,
	}
}

func (fh *facultyHandlers) GetFaculties() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "faculty.GetFaculties")
		defer span.Finish()

		faculties, err := fh.fu.GetFaculties(ctx)
		if err != nil {
			logger.ErrResponseWithLog(c, fh.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, faculties)
	}
}

func (fh *facultyHandlers) CreateFaculty() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "faculty.CreateFaculty")
		defer span.Finish()

		var faculty models.Faculty
		if err := c.ShouldBindJSON(&faculty); err != nil {
			logger.ErrResponseWithLog(c, fh.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		err := fh.fu.CreateFaculty(ctx, &faculty)
		if err != nil {
			logger.ErrResponseWithLog(c, fh.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, faculty)
	}
}

func (fh *facultyHandlers) UpdateFaculty() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "faculty.UpdateFaculty")
		defer span.Finish()

		var faculty models.Faculty
		if err := c.ShouldBindJSON(&faculty); err != nil {
			logger.ErrResponseWithLog(c, fh.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		err := fh.fu.UpdateFaculty(ctx, &faculty)
		if err != nil {
			logger.ErrResponseWithLog(c, fh.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, faculty)
	}
}

func (s *facultyHandlers) DeleteFaculty() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "faculty.DeleteFaculty")
		defer span.Finish()

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			logger.ErrResponseWithLog(c, s.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		err = s.fu.DeleteFaculty(ctx, uint(id))
		if err != nil {
			logger.ErrResponseWithLog(c, s.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, nil)
	}
}
