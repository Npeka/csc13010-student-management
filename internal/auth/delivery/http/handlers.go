package http

import (
	"net/http"

	"github.com/csc13010-student-management/internal/auth"
	"github.com/csc13010-student-management/internal/auth/dtos"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/csc13010-student-management/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

type authHandlers struct {
	au auth.IAuthUsecase
	lg *logger.LoggerZap
}

func NewAuthHandlers(
	au auth.IAuthUsecase,
	lg *logger.LoggerZap,
) auth.IAuthHandlers {
	return &authHandlers{
		au: au,
		lg: lg,
	}
}

func (a *authHandlers) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "auth.Register")
		defer span.Finish()

		var registerReq dtos.UserRegisterRequestDTO
		if err := c.ShouldBindJSON(&registerReq); err != nil {
			logger.ErrResponseWithLog(c, a.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		err := a.au.Register(ctx, &registerReq)
		if err != nil {
			logger.ErrResponseWithLog(c, a.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.UserRegiterSuccess, nil)
	}
}

func (a *authHandlers) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "auth.Login")
		defer span.Finish()

		var loginReq dtos.UserLoginRequestDTO
		if err := c.ShouldBindJSON(&loginReq); err != nil {
			logger.ErrResponseWithLog(c, a.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		loginRes, err := a.au.Login(ctx, &loginReq)
		if err != nil {
			logger.ErrResponseWithLog(c, a.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, loginRes)
	}
}

func (a *authHandlers) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func (a *authHandlers) Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
