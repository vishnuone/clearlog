package clearlog

import (
	"log"
	"runtime"
	"strconv"
)


func Info(message string) {
	log.Print("\tinfo\t" + message)
}

func Error(error error, msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	log.Print("\terror\t" + file + ":" + strconv.Itoa(line) + "\t" + error.Error() + "\t" + msg)
}
