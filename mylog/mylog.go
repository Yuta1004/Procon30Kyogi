package mylog

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Info : タイプInfoのログを出力する
func Info(fmt string, args ...interface{}) {
	header := "[INFO] "
	footer := "\x1b[0m\n"
	outlog(header, fmt, footer, args...)
}

// Error : タイプErrorのログを出力する
func Error(fmt string, args ...interface{}) {
	header := "\x1b[31m[ERROR] "
	footer := "\x1b[0m\n"
	outlog(header, fmt, footer, args...)
}

// Notify : タイプNotifyのログを出力する
func Notify(fmt string, args ...interface{}) {
	header := "\x1b[32m[NOTIFY] "
	footer := "\x1b[0m\n"
	outlog(header, fmt, footer, args...)
}

// Warning : タイプWarningのログを表示する
func Warning(fmt string, args ...interface{}) {
	header := "\x1b[33m[WARNING] "
	footer := "\x1b[0m\n"
	outlog(header, fmt, footer, args...)
}

func outlog(header, fmtStr, footer string, args ...interface{}) {
	l := log.New(os.Stdout, "\x1b[5A\x1b[G"+time.Now().Format("2006/01/02 15:05:04.000 "), 0)
	if len(args) > 0 {
		l.Printf(header+fmtStr+footer, args...)
	} else {
		l.Printf(header + fmtStr + footer)
	}
	fmt.Printf("\x1b[G\x1b[K\n\x1b[K\n\x1b[K\n\x1b[K\n\x1b[K\n\x1b[K\x1b[G>> ")
}
