package services

import (
	"LinkUp_Update/var/logs"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Get(c *gin.Context, db *sqlx.DB, ctx context.Context) *Service {
	return &Service{
		logs.Get(),
		c,
		db,
		ctx,
	}
}

type identification interface {
	Auth()
	Register()
	Logout()
}

type Service struct {
	l   *logs.Logging
	c   *gin.Context
	db  *sqlx.DB
	ctx context.Context
}
