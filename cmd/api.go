package cmd

import (
	"LinkUp_Update/internal/routing"
	"github.com/gin-gonic/gin"
)

func Run() {
	var s = &MyServer{}
	r := gin.New()
	routing.StartRouting(r)
	s.StartServer(r)
}
