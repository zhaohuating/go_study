package middleware

import (
	"fmt"
	"net/http"
	"os"
	"task4/config"
	"time"

	"github.com/gin-gonic/gin"
	rotatelog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	filePath := config.Cfg.Log.OutputPath
	logLevel := config.Cfg.Log.Level
	logger := logrus.New()
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		fmt.Printf("parse log level error, The default level has been set to: info。%s\n", err.Error())
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)
	age := config.Cfg.Log.MaxAge
	if age <= 0 {
		age = 1
	}
	maxAge := age * 24
	maxAgeHour := time.Duration(maxAge) * time.Hour
	logWriter, _ := rotatelog.New(
		filePath+"%Y-%m-%d.log",
		rotatelog.WithMaxAge(maxAgeHour),
		rotatelog.WithRotationTime(24*time.Hour),
		rotatelog.WithLinkName("latest_log.log"),
	)

	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.PanicLevel: logWriter,
		logrus.TraceLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
	}

	lfHook := lfshook.NewHook(writerMap, &logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
	})
	logger.AddHook(lfHook)
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := fmt.Sprintf("%d ms", end.Sub(start).Milliseconds())
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataLength := c.Writer.Size()
		path := c.Request.RequestURI
		method := c.Request.Method
		statusCode := c.Writer.Status()
		entry := logger.WithFields(logrus.Fields{
			"hostname":   hostname,
			"statusCode": statusCode,
			"latency":    latency,
			"clientIP":   clientIP,
			"method":     method,
			"path":       path,
			"userAgent":  userAgent,
			"dataLength": dataLength,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}

		if statusCode >= http.StatusInternalServerError {
			entry.Error()
		} else if statusCode >= http.StatusBadRequest {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
