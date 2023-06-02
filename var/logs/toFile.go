package logs

import (
	"fmt"
	"os"
)

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
