package server

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type PromMiddleware struct {
	requests *prometheus.HistogramVec
}

func NewPromMiddleware() PromMiddleware {

	h := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "service",
		Name:      "request-metrics",
		Help:      "HTTP requests partionied by code, http, method, and path",
		Buckets:   []float64{0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 5},
	}, []string{"code", "method", "path"})

	return PromMiddleware{
		requests: h,
	}
}

func (m *PromMiddleware) CountRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		dur := float64(time.Since(start))

		m.requests.WithLabelValues("status", r.Method, r.URL.Path).Observe(dur)
	})
}
