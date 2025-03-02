package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/csc13010-student-management/pkg/response"
	"github.com/csc13010-student-management/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
)

type MiddlewareManager struct {
	e *casbin.Enforcer
}

func NewMiddlewareManager(e *casbin.Enforcer) *MiddlewareManager {
	return &MiddlewareManager{
		e: e,
	}
}

func (mw *MiddlewareManager) RBAC(obj, act string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := jwt.ExtractBearerToken(c.Request)
		if token == "" {
			response.Error(c, response.ErrMissingToken)
			c.Abort()
			return
		}

		claims, err := jwt.ValidateJWTToken(token)
		if err != nil {
			response.Error(c, response.ErrInvalidToken)
			c.Abort()
			return
		}

		username := claims.Username
		if username == "" {
			response.Error(c, response.ErrInvalidToken)
			c.Abort()
			return
		}

		obj := c.Request.URL.Path
		act := c.Request.Method

		ok, err := mw.e.Enforce(username, obj, act)
		if err != nil {
			response.Error(c, response.ErrCheckPermission)
			c.Abort()
			return
		}

		if !ok {
			response.Error(c, response.ErrPermissionDenied)
			c.Abort()
			return
		}

		c.Next()
	}
}
