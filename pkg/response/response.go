package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code].message,
		Data:    data,
	})
}

func Error(c *gin.Context, code int) {
	c.JSON(msg[code].httpCode, ResponseData{
		Code:    code,
		Message: msg[code].message,
		Data:    nil,
	})
}
