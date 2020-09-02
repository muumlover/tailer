package logger

import (
	"log"
	"os"
)

const (
	LevelDebug = 1 << iota
	LevelTrace
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

var (
	debugLogger   = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags|log.Llongfile)
	traceLogger   = log.New(os.Stdout, "[TRACE] ", log.LstdFlags|log.Llongfile)
	infoLogger    = log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Llongfile)
	warningLogger = log.New(os.Stdout, "[WARNING] ", log.LstdFlags|log.Llongfile)
	errorLogger   = log.New(os.Stdout, "[ERROR] ", log.LstdFlags|log.Llongfile)
	fatalLogger   = log.New(os.Stdout, "[FATAL] ", log.LstdFlags|log.Llongfile)
)

func logger(level int) func(v ...interface{}) {
	switch level {
	case LevelDebug:
		return debugLogger.Println
	case LevelTrace:
		return traceLogger.Println
	case LevelInfo:
		return infoLogger.Println
	case LevelWarning:
		return warningLogger.Println
	case LevelError:
		return errorLogger.Println
	case LevelFatal:
		return fatalLogger.Println
	}
	return nil
}

var (
	Debug   = logger(LevelDebug)
	Trace   = logger(LevelTrace)
	Info    = logger(LevelInfo)
	Warning = logger(LevelWarning)
	Error   = logger(LevelError)
	Fatal   = logger(LevelFatal)
)
