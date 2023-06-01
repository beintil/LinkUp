package middleware

import (
	cs "LinkUp_Update/constans"
	mycookies "LinkUp_Update/internal/cookie"
	"LinkUp_Update/internal/database"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// CheckCookieAndRedirect A method of Middleware that checks if a cookie
// is present in the request, and if not, redirects to a /auth
func (mw *Middleware) CheckCookieAndRedirect() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !mycookies.CookiesExist(c) {
			c.Redirect(http.StatusSeeOther, cs.Authorization)
			return
		}
	}
}

func (mw *Middleware) HandlerVisitors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				mw.l.LogApi(rec)
			}
		}()
		localId := c.Param("id")
		idInt, err := strconv.ParseInt(localId, 10, 32)

		if localId != "" && err == nil && idInt <= 999999999 && localId[:1] != "0" {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			db := database.Init(ctx).GetConn()
			defer database.Close(db)

			var existUserWithId bool
			err = db.Get(&existUserWithId, "SELECT EXISTS(SELECT 1 FROM users WHERE local_id = $1)", localId)
			if err != nil {
				mw.l.LogApi(err)
				c.Status(http.StatusInternalServerError)
				return
			}
			if !existUserWithId {
				c.Redirect(http.StatusSeeOther, cs.Home)
				return
			}
			var owner bool

			err = db.Get(&owner, "SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 AND local_id = $2)", mycookies.DecodeIdFromCookie(c), localId)
			if err != nil {
				mw.l.LogApi(err)
				c.Status(http.StatusInternalServerError)
				return
			}
			if owner {
				mw.redirectTo(c)
				return
			}
			c.Next()
			return
		}
		c.Redirect(http.StatusSeeOther, cs.Home)
	}
}

func (mw *Middleware) redirectTo(c *gin.Context) {
	switch c.Request.URL.Path {
	case cs.UrlWithoutId(cs.Home, c.Request.URL.Path):
		c.Redirect(http.StatusSeeOther, cs.Home)
	}
}
