package friends

import (
	mycookies "LinkUp_Update/internal/cookie"
	"LinkUp_Update/internal/database"
	"LinkUp_Update/internal/friends/service"
	"LinkUp_Update/var/logs"
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func Get(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	db := database.Init(ctx).GetConn()
	defer func() {
		database.Close(db)
		if rec := recover(); rec != nil {
			logs.Get().LogApi(rec)
		}
	}()
	service.GetService(db, c).Get("SELECT friends_id FROM users WHERE id = $1", mycookies.DecodeIdFromCookie(c))
}
