package routing

import (
	"LinkUp_Update/var/logs"
	"github.com/gin-gonic/gin"
)

func StartRouting(router *gin.Engine) {
	routing(&routers{
		router,
		logs.Get(),
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
	r.identification()
	r.profile()
	r.friends()
	r.search()
	r.user()
}
