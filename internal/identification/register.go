package identification

import (
	"LinkUp_Update/internal/database"
	"LinkUp_Update/internal/identification/services"
	"LinkUp_Update/var/logs"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Registration(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		services.FormHandler(c)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	db := database.Init(ctx).GetConn()
	defer func() {
		database.Close(db)
		if rec := recover(); rec != nil {
			logs.Get().LogApi(rec)
		}
	}()
	services.Get(c, db, ctx).Register()
}
