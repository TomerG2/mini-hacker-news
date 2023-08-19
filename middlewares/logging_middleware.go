package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		startTime := time.Now()
		logrus.Infof("[START] HTTP REQUEST [route=%s] [method=%s]", reqUri, reqMethod)

		ctx.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		statusCode := ctx.Writer.Status()
		logrus.Infof("[FINISH] HTTP REQUEST [route=%s] [method=%s] [status=%d] [latency=%d]", reqUri, reqMethod, statusCode, latencyTime.Milliseconds())
		ctx.Next()
	}
}
