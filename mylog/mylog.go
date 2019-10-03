package mylog

import (
	"log"
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

func outlog(header, fmt, footer string, args ...interface{}) {
	if len(args) > 0 {
		log.Printf(header+fmt+footer, args...)
	} else {
		log.Printf(header + fmt + footer)
	}
}
