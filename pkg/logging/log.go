package logging

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"simple_blog/pkg/setting"
)

type Level int

var F *os.File
var logger *log.Logger

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func SetupLog(config setting.Setting) {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)
	logger = log.New(F, config.DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}


