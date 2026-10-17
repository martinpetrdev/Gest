package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	scopeName string
}

func New(scopeName string) *Logger {
	return &Logger{
		scopeName: scopeName,
	}
}

func (l *Logger) Log(line string) {
	fmt.Printf("[%s] %-8d - %s %5s [%s] %s\n", "Gest", os.Getpid(), time.Now().UTC().Format("02/01/2006 15:04:05"), "LOG", l.scopeName, line)
}
