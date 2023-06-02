package logs

import (
	"LinkUp_Update/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func (l *Logging) LogHttpRequest(c *gin.Context) {
	if config.Get("LOGGING_HTTP").ToBool() {
		l.mx.RLock()
		defer l.mx.RUnlock()
		r := c.Request

		logHTTP := fmt.Sprintf("\nMETHOD: %s, URL: %s, ID_USER: %s\nHOST: %s\nHTTP: %s\nUSER-AGENT: %s", r.Method, r.URL.Path, mux.Vars(r)["id"], r.Host, r.Proto, r.UserAgent())
		fmt.Println(logHTTP)

		if config.Get("LOGGING_HTTP").ToBool() {
			l.loggingToFile("./var/logs/logs_http.txt", logHTTP)
		}
	}
}
