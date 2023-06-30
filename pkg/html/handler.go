package html

import (
	"LinkUp_Update/var/logs"
	"github.com/gin-gonic/gin"
	"html/template"
)

func HandlerWithEntity(c *gin.Context, fileName string, entity any) {
	tmpl, err := template.ParseFiles("./pkg/html/" + fileName)
	if err != nil {
		logs.Get().LogApi(err)
		return
	}

	err = tmpl.Execute(c.Writer, entity)
	if err != nil {
		logs.Get().LogApi(err)
	}
}

func HandlerNotEntity(c *gin.Context, fileName string) {
	tmpl, err := template.ParseFiles("./pkg/html/" + fileName)
	if err != nil {
		logs.Get().LogApi(err)
		return
	}
	err = tmpl.Execute(c.Writer, nil)
	if err != nil {
		logs.Get().LogApi(err)
	}
}
