package logger

import (
	"log"
	"os"
)

const (
	LevelTrace = 1 << iota
	LevelDebug
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

func logger(level int) func(format string, v ...interface{}) {
	switch level {
	case LevelDebug:
		return debugLogger.Printf
	case LevelTrace:
		return traceLogger.Printf
	case LevelInfo:
		return infoLogger.Printf
	case LevelWarning:
		return warningLogger.Printf
	case LevelError:
		return errorLogger.Printf
	case LevelFatal:
		return fatalLogger.Printf
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
