package friends

import (
	cs "LinkUp_Update/constans"
	"LinkUp_Update/internal/database"
	"LinkUp_Update/internal/friends/service"
	"LinkUp_Update/var/logs"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Add(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	db := database.Init(ctx).GetConn()
	defer func() {
		database.Close(db)
		if rec := recover(); rec != nil {
			logs.Get().LogApi(rec)
		}
	}()
	service.GetService(db, c).Add()
	c.Redirect(http.StatusSeeOther, cs.Conversion(cs.GetFriends))
}
