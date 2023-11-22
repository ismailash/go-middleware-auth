package middleware

import (
	"log"
	"time"

	"enigmacamp.com/be-enigma-laundry/utils/common"
	modelutil "enigmacamp.com/be-enigma-laundry/utils/model_util"
	"github.com/gin-gonic/gin"
)

type LogMiddleware interface {
	LogRequest() gin.HandlerFunc
}

type logMiddleware struct {
	logService common.MyLogger
}

func (l *logMiddleware) LogRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := l.logService.InitializeLogger(); err != nil {
			log.Fatal("err: ", err.Error())
		}
		t := time.Now()

		logString := modelutil.RequestLog{
			AccessTime: t,
			Latency:    time.Since(t),
			ClientIP:   ctx.ClientIP(),
			Method:     ctx.Request.Method,
			Code:       ctx.Writer.Status(),
			Path:       ctx.Request.URL.Path,
			UserAgent:  ctx.Request.UserAgent(),
		}

		switch {
		case ctx.Writer.Status() >= 500:
			l.logService.LogFatal(logString)
		case ctx.Writer.Status() >= 400:
			l.logService.LogWarn(logString)
		default:
			l.logService.LogInfo(logString)
		}
	}
}

func NewLogMiddleware(logService common.MyLogger) LogMiddleware {
	return &logMiddleware{logService: logService}
}
