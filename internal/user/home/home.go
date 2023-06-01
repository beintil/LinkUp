package home

import (
	"LinkUp_Update/var/logs"
	"github.com/gin-gonic/gin"
	"html/template"
)

func Home(c *gin.Context) {
	formHandler(c)
}

func formHandler(c *gin.Context) {
	tmpl, err := template.ParseFiles("./internal/user/home/html/home.html")
	if err != nil {
		logs.Get().LogApi(err)
		return
	}
	err = tmpl.Execute(c.Writer, nil)
	if err != nil {
		logs.Get().LogApi(err)
		return
	}
}
