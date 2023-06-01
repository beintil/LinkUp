package logs

import (
	"net/http"
	"sync"
)

type logging interface {
	LogApi(err error, f ...func())
	LogHttpRequest(r *http.Request)
	loggingToFile(file string, logMsg string)
}

type Logging struct {
	mx   *sync.RWMutex
	err  any
	line int
	file string
}
