package service

import (
	"LinkUp_Update/var/logs"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func GetService(db *sqlx.DB, c *gin.Context) *Service {
	return &Service{
		db,
		logs.Get(),
		c,
	}
}

type searchUsers interface {
}

type Service struct {
	db *sqlx.DB
	l  *logs.Logging
	c  *gin.Context
}
