package logs

import (
	"LinkUp_Update/config"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

func Get() *Logging {
	return &Logging{
		mx: &sync.RWMutex{},
	}
}

func (l *Logging) LogApi(err any, f ...func()) {
	l.mx.RLock()
	defer l.mx.RUnlock()

	defer func() {
		if rec := recover(); rec != nil {
			log.Fatal(rec)
		}
	}()

	red := color.New(color.FgHiRed).SprintFunc()
	green := color.New(color.FgHiGreen).SprintFunc()
	while := color.New(color.BgWhite).SprintFunc()
	var ok bool
	l.err = err
	_, l.file, l.line, ok = runtime.Caller(1)
	if !ok {
		l.file = "unknown"
		l.line = 0
	}

	if config.Get("LOGGING_API").ToBool() {
		l.loggingToFile("./var/logs/logs_api.txt", fmt.Sprintf("\n%v: %v\nFile: %v\nLine: %v\n\n", time.Now().Format("2006-01-02 15:04:05"), l.err, l.file, l.line))
	}

	fmt.Print(fmt.Sprintf("\n%v: %v\n"+while(green("File:"))+"%v\n"+while(green("Line:"))+"%v\n\n", while(green(time.Now().Format("2006-01-02 15:04:05"))), red(l.err), green(l.file), green(l.line)))

	if f != nil {
		f[0]()
	}
}

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

func (l *Logging) loggingToFile(file string, logMsg string) {
	logFile, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer func() {
		err = logFile.Close()
		if err != nil {
			panic(err)
		}
	}()

	if _, err = logFile.WriteString(logMsg); err != nil {
		fmt.Println("Failed to write to log file:", err)
	}
}
