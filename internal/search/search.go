package search

import (
	"LinkUp_Update/internal/database"
	"LinkUp_Update/internal/search/service"
	"LinkUp_Update/pkg/html"
	"LinkUp_Update/var/logs"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Search(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		html.HandlerWithEntity(c, "search.html", nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var db = database.Init(ctx).GetConn()
	defer func() {
		database.Close(db)
		if rec := recover(); rec != nil {
			logs.Get().LogApi(rec)
		}
	}()
	service.GetService(db, c).Search()
}
