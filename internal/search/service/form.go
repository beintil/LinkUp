package service

import (
	"LinkUp_Update/var/logs"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func DataFormHandler(c *gin.Context, users *[]search) {
	tmpl, err := template.ParseFiles("./internal/search/html/search.html")
	if err != nil {
		logs.Get().LogApi(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	if users != nil {
		err = tmpl.Execute(c.Writer, *users)
	} else {
		err = tmpl.Execute(c.Writer, nil)
	}

	if err != nil {
		logs.Get().LogApi(err)
		c.Status(http.StatusInternalServerError)
	}
}
