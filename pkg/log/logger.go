package log

import "github.com/op/go-logging"

const format = `%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s}%{id:03x}%{color:reset} %{message}`

var (
	log = logging.MustGetLogger("monito")

	formatter = logging.MustStringFormatter(format)
)

//Debug helper function for debugging
func Debug(format string, args ...interface{}) { log.Debug(format, args...) }

//Info helper function to test out funcionality
func Info(format string, args ...interface{}) { log.Info(format, args...) }

//Warning helper function for warnings thrown
func Warning(format string, args ...interface{}) { log.Warning(format, args...) }

//Critical errors thrown that should be logged.
func Critical(format string, args ...interface{}) { log.Critical(format, args...) }

//Panic helper function
func Panic(format string, args ...interface{}) { log.Panicf(format, args...) }
