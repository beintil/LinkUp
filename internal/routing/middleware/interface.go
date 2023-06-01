package middleware

import (
	"LinkUp_Update/var/logs"
	"github.com/gin-gonic/gin"
)

func Get() *Middleware {
	return &Middleware{
		logs.Get(),
	}
}

type middlewares interface {
	CheckCookieAndRedirect() gin.HandlerFunc
	HandlerVisitors() gin.HandlerFunc
	redirectTo(c *gin.Context)
}

type Middleware struct {
	l *logs.Logging
}
