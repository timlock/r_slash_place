package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		slog.Info(fmt.Sprintf("%s %s %s", req.Method, req.RequestURI, time.Since(start)))
	})
}

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				slog.Error((string(debug.Stack())))
			}
		}()
		next.ServeHTTP(w, req)
	})
}
