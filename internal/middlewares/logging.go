package middlewares

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func SlogMiddleware() gin.HandlerFunc {
	logger := slog.Default()
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()

		logger.Info("HTTP Request",
			"path", path,
			"raw_query", raw,
			"status", statusCode,
			"method", c.Request.Method,
			"latency", latency.String(),
			"client_ip", c.ClientIP(),
		)
	}
}
