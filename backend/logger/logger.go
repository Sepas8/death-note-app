package logger

import (
	"log"
	"net/http"
	"time"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{log.New(log.Writer(), "", 0)}
}

func (l *Logger) RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		log.Printf("\033[44m %s \033[0m | PATH: \033[33m\"%s\"\033[0m | TIME: %v", r.Method, r.URL.Path, now.Format("2006-01-02 15:04:05"))
		next.ServeHTTP(w, r)
	})
}

func (l *Logger) Info(status int, path string, start time.Time) {
	l.Printf("\033[42m %d \033[0m | PATH: \033[33m\"%s\"\033[0m | DURATION: \033[42m %v \033[0m", status, path, time.Since(start))
}

func (l *Logger) Error(status int, path string, err error) {
	l.Printf("\033[41m %d \033[0m | PATH: \033[33m\"%s\"\033[0m | ERROR: \033[31m %v \033[0m", status, path, err)
}

func (l *Logger) Fatal(err error) {
	l.Fatalf("\033[41m FATAL \033[0m | ERROR: \033[31m %v \033[0m", err)
}
