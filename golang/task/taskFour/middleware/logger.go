package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

func init() {
	// 设置Viper读取配置文件
	viper.SetConfigName("config") // 配置文件名（不带扩展名）
	viper.SetConfigType("yaml")   // 如果需要指定配置文件类型
	viper.AddConfigPath(".")      // 查找当前目录下的配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	// 设置日志格式为json格式，也可以选择text格式
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint:     true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 设置输出到标准输出（可以改为文件或其他目的地）
	logrus.SetOutput(os.Stdout)
	// 设置日志级别，默认是info级别
	logrus.SetLevel(logrus.InfoLevel)
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 记录请求开始时间
		start := time.Now()
		ctx.Next()
		// 获取响应状态码
		statusCode := ctx.Writer.Status()
		// 构造日志条目
		entry := logrus.WithFields(logrus.Fields{
			"method":  ctx.Request.Method,
			"path":    ctx.Request.URL.Path,
			"status":  statusCode,
			"latency": fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
		})
		// 根据状态码确定日志级别
		if statusCode != 200 {
			entry.Warn("error")
		} else {
			entry.Info("success")
		}
	}
}
