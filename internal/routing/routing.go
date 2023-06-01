package routing

import (
	"LinkUp_Update/internal/routing/middleware"
	"LinkUp_Update/var/logs"
	"github.com/gin-gonic/gin"
)

func StartRouting(router *gin.Engine) {
	routing(&routers{
		router,
		logs.Get(),
		middleware.Get(),
	})
}

func routing(r *routers) {
	r.eng.NoRoute(func(c *gin.Context) {
		_, err := c.Writer.Write([]byte("the page not exist"))
		if err != nil {
			logs.Get().LogApi(err)
		}
	})
	r.eng.Use(r.l.LogHttpRequest)

	// for routers not requiring authorization
	r.identification()

	// for routers requiring authorization
	r.eng.Use(r.mw.CheckCookieAndRedirect())
	r.profile()
	r.friends()
	r.search()
	r.user()
	r.visit()
}
