package clearlog

import (
	"log"
	"runtime"
	"strconv"
)


func Info(message string) {
	log.Print(message)
}

func Error(error error, msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	log.Print(file + ":" + strconv.Itoa(line) + "\n" + error.Error() + "\n" + msg)
}
