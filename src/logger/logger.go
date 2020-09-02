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
	debugLogger   = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags|log.Lshortfile)
	traceLogger   = log.New(os.Stdout, "[TRACE] ", log.LstdFlags|log.Lshortfile)
	infoLogger    = log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "[WARNING] ", log.LstdFlags|log.Lshortfile)
	errorLogger   = log.New(os.Stdout, "[ERROR] ", log.LstdFlags|log.Lshortfile)
	fatalLogger   = log.New(os.Stdout, "[FATAL] ", log.LstdFlags|log.Lshortfile)
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
