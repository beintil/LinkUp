package home

import (
	"LinkUp_Update/var/logs"
	"github.com/gin-gonic/gin"
	"html/template"
)

func Visit(c *gin.Context) {
	visitFormHandler(c)
}

func visitFormHandler(c *gin.Context) {
	tmpl, err := template.ParseFiles("./internal/user/home/html/home_visit.html")
	if err != nil {
		logs.Get().LogApi(err)
		return
	}
	err = tmpl.Execute(c.Writer, c.Param("id"))
	if err != nil {
		logs.Get().LogApi(err)
		return
	}
}
