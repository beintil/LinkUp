package logs

import (
	"LinkUp_Update/config"
	"fmt"
	"github.com/fatih/color"
	"log"
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
