package home

import (
	cs "LinkUp_Update/constans"
	"LinkUp_Update/var/logs"
	"github.com/gin-gonic/gin"
	"html/template"
)

func Home(c *gin.Context) {}

func FormHandler(c *gin.Context) {
	var url string
	if cs.UrlWithoutId(cs.Home, c.Request.URL.Path) == c.Request.URL.Path {
		url = "./internal/user/home/html/home_rights_is_guest.html"
	} else {
		url = "./internal/user/home/html/home.html"
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
