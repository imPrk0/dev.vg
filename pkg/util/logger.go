package util

import (
	"fmt"
	"github.com/fatih/color"
	"sync"
	"time"
)

type Logger struct {
	level int
	mu    sync.Mutex
}

const (
	LevelError = iota
	LevelWarning
	LevelInformational
	LevelDebug
)

var (
	GlobalLogger *Logger
	Level        = LevelDebug
	colors       = map[string]func(a ...interface{}) string{
		"Warning": color.New(color.FgYellow).Add(color.Bold).SprintFunc(),
		"Panic":   color.New(color.BgRed).Add(color.Bold).SprintFunc(),
		"Error":   color.New(color.FgRed).Add(color.Bold).SprintFunc(),
		"Info":    color.New(color.FgCyan).Add(color.Bold).SprintFunc(),
		"Debug":   color.New(color.FgWhite).Add(color.Bold).SprintFunc(),
	}
	spaces = map[string]string{
		"Warning": "",
		"Panic":   "  ",
		"Error":   "  ",
		"Info":    "   ",
		"Debug":   "  ",
	}
)

func (l *Logger) Println(prefix string, msg string) {
	c := color.New()

	l.mu.Lock()
	defer l.mu.Unlock()

	_, _ = c.Printf(
		"%s%s %s %s\n",
		colors[prefix]("["+prefix+"]"),
		spaces[prefix],
		time.Now().Format("2006-01-02 15:04:05"),
		msg,
	)
}

func Log() *Logger {
	if nil == GlobalLogger {
		l := Logger{
			level: Level,
		}
		GlobalLogger = &l
	}
	return GlobalLogger
}

// ????
//func (l *Logger) Success(format string, v ...interface{}){
//}

func (l *Logger) Info(format string, v ...interface{}) {
	if LevelInformational > l.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	l.Println("Info", msg)
}

func (l *Logger) Warning(format string, v ...interface{}) {
	if LevelWarning > l.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	l.Println("Warning", msg)
}

func (l *Logger) Error(format string, v ...interface{}) {
	if LevelError > l.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	l.Println("Error", msg)
}

func (l *Logger) Panic(format string, v ...interface{}) {
	if LevelError > l.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	l.Println("Panic", msg)
	panic(msg)
}

func (l *Logger) Debug(format string, v ...interface{}) {
	if LevelDebug > l.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	l.Println("Debug", msg)
}

func BuildLogger(level string) {
	intLevel := LevelError
	switch level {
	case "error":
		intLevel = LevelError
	case "warning":
		intLevel = LevelWarning
	case "info":
		intLevel = LevelInformational
	case "debug":
		intLevel = LevelDebug
	}
	l := Logger{
		level: intLevel,
	}
	GlobalLogger = &l
}
