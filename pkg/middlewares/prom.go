package middlewares

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// https://www.robustperception.io/prometheus-middleware-for-gorilla-mux

var (
	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "ports_http_duration_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})
)

// PrometheusMiddleware implements mux.MiddlewareFunc.
func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, err := route.GetPathTemplate()
		if err != nil {
			next.ServeHTTP(w, r)
		}
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		next.ServeHTTP(w, r)
		timer.ObserveDuration()
	})
}
