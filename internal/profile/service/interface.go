package service

import (
	"LinkUp_Update/var/logs"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func GetService(db *sqlx.DB, c *gin.Context) *Service {
	return &Service{
		logs.Get(),
		db,
		c,
	}
}

type profile interface {
	Get()
	Edit()
}

type Service struct {
	l  *logs.Logging
	db *sqlx.DB
	c  *gin.Context
}
