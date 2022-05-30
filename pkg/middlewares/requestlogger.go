package middlewares

import (
	"log"
	"net/http"
	"time"
)

func RequestLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		log.Printf("ip: '%s' - method: '%s' - path: '%s' - ms: '%d'\n", r.RemoteAddr, r.Method, r.URL, time.Since(start).Milliseconds())
	})
}
