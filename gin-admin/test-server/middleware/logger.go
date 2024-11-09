package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	logClient := logrus.New()
	//禁止logrus的输出
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err:", err)
	}
	logClient.Out = src
	logClient.SetLevel(logrus.DebugLevel)
	apiLogPath := "api.log"
	rotatelogs.New(
		apiLogPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(apiLogPath),       // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logClient.Writer(),
		logrus.FatalLevel: logClient.Writer(),
		logrus.DebugLevel: logClient.Writer(),
	}
	lfsHook := lfshook.NewHook(writerMap, &logrus.JSONFormatter{})
	logClient.AddHook(lfsHook)

	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		// 执行时间
		latency := end.Sub(start)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqPath := c.Request.URL.Path
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()

		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		requestParams := buf[:n]

		logClient.Info(
			"| %3d | %13v | %15s | %s  %s |%s|",
			statusCode,
			latency,
			clientIP,
			reqMethod,
			reqPath,
			requestParams,
		)
	}
}
