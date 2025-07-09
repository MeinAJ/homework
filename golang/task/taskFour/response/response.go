package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SuccessCode = 0
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ReturnSuccess 成功返回
func ReturnSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Success(data))
}

// ReturnError 成功返回
func ReturnError(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(http.StatusOK, Error(code, message, data))
}

// Success 成功返回
func Success(data interface{}) *Response {
	return &Response{
		Code:    SuccessCode,
		Message: "success",
		Data:    data,
	}
}

// Custom 自定义返回
func Custom(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// Error 错误返回
func Error(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
