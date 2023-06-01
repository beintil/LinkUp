package profile

import (
	"LinkUp_Update/internal/database"
	"LinkUp_Update/internal/profile/service"
	"LinkUp_Update/var/logs"
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func Get(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var db = database.Init(ctx).GetConn()
	defer func() {
		database.Close(db)
		if rec := recover(); rec != nil {
			logs.Get().LogApi(rec)
		}
	}()
	service.GetService(db, c).Get()
}
