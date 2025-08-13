package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	msg  string
	time time.Time
	err  error
}

func NewLog(msg string, err error) Logger {
	return Logger{msg: msg, time: time.Now(), err: err}
}

func AsLogger() chan Logger {
	logchan := make(chan Logger, 50)
	go func() {
		for v := range logchan {
			fmt.Println(v)
		}
	}()
	return logchan
}
