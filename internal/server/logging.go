package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type ZapMiddleware struct {
	logger *zap.Logger
}

func NewZapMiddleware(l *zap.Logger) ZapMiddleware {

	return ZapMiddleware{
		logger: l,
	}
}

func (z *ZapMiddleware) LogRequest(h http.Handler) http.Handler {

	f := func(w http.ResponseWriter, r *http.Request) {

		wr := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()

		h.ServeHTTP(w, r)

		dur := time.Since(start)
		z.logger.With(
			zap.Int("status", wr.Status()),
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Duration("duration", dur),
		)
	}

	return http.HandlerFunc(f)
}
