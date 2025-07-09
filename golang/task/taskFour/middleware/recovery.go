package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang-homework/task/taskFour/response"
)

// RecoveryMiddleware 全局异常捕获中间件
func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			// 捕获异常，判断是否是*response.Response类型，是则返回错误信息，否则返回500错误
			if err := recover(); err != nil {
				if _, ok := err.(*response.Response); ok {
					errorResponse := err.(*response.Response)
					// 返回统一错误格式
					logrus.Error(errorResponse.Message)
					response.ReturnError(ctx, errorResponse.Code, errorResponse.Message, errorResponse.Data)
					ctx.Abort()
				} else {
					// 返回统一错误格式
					logrus.Error("Internal Server Error")
					response.ReturnError(ctx, 500, "Internal Server Error", nil)
					ctx.Abort()
				}
			}
		}()
		ctx.Next()
	}
}
