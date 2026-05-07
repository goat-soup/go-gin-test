package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F                  *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	logger             *log.Logger
	logPrefiex         = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

func init() {
	filepath := getLogFileFullPath()                  // 获取日志文件路径
	F = OpenLogFile(filepath)                         // 打开日志文件
	logger = log.New(F, DefaultPrefix, log.LstdFlags) // 创建日志记录器
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Print(v...)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Print(v...)
}
func Warn(v ...interface{}) {
	setPrefix(WARN)
	logger.Print(v...)
}
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Print(v...)
}
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Print(v...)
}
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefiex = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefiex = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefiex)
}
