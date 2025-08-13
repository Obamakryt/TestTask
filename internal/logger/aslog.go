package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	msg  string
	time time.Time
	err  string
}

func NewLog(msg string, err string) Logger {
	return Logger{msg: msg, time: time.Now(), err: err}
}

func AsLogger() chan Logger {
	logchan := make(chan Logger, 50)
	go func() {
		for v := range logchan {
			fmt.Printf("%s,%s, Error=%s\n",
				v.time.Format("2006-01-02 15:04:05"),
				v.msg,
				v.err,
			)
		}
	}()
	return logchan
}
