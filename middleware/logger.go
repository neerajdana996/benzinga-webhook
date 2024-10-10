package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// RequestLogger logs requests and startup initialization
func RequestLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        logrus.Infof("Incoming request: %s %s", c.Request.Method, c.Request.URL)
        c.Next()
    }
}
