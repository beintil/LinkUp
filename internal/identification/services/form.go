package services

import (
	"LinkUp_Update/var/logs"
	"github.com/gin-gonic/gin"
	"html/template"
)

func FormHandler(c *gin.Context) {
	var url string
	switch c.Request.URL.Path {
	case "/register":
		url = "./internal/identification/html/registration.html"
	case "/auth":
		url = "./internal/identification/html/auth.html"
	}
	tmpl, err := template.ParseFiles(url)
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
