package middleware

import (
	cs "LinkUp_Update/constans"
	mycookies "LinkUp_Update/internal/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Middleware struct {
}

// CheckCookieAndRedirect A method of Middleware that checks if a cookie
// is present in the request, and if not, redirects to a /auth
func (mw *Middleware) CheckCookieAndRedirect() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !mycookies.CookiesExist(c) {
			c.Redirect(http.StatusSeeOther, cs.Conversion(cs.Authorization))
			return
		}
	}
}
