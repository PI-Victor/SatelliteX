package log

import "github.com/op/go-logging"

const format = `%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s}%{id:03x}%{color:reset} %{message}`

var (
	log = logging.MustGetLogger("oinc")

	formatter = logging.MustStringFormatter(format)
)

//Info helper function to test out funcionality
func Info(format string, args ...interface{}) { log.Debug(format, args...) }

//Panic helper function
func Panic(format string, args ...interface{}) { log.Panicf(format, args...) }
