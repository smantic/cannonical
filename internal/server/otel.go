package server

import (
	"net/http"

	"go.opentelemetry.io/otel"
)

func AddTrace(h http.Handler) http.Handler {

	f := func(w http.ResponseWriter, r *http.Request) {
		ctx, span := otel.Tracer("server").Start(r.Context(), "AddTrace")
		trace := r.WithContext(ctx)

		h.ServeHTTP(w, trace)
		span.End()
	}
	return http.HandlerFunc(f)
}
