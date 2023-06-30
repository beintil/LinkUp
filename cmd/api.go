package cmd

import (
	"LinkUp_Update/internal/routing"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.New()
	routing.StartRouting(r)
	StartServer(r)
}
