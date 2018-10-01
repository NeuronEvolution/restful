package rest

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		l := zap.L().Named("request").
			With(zap.String("url", r.RequestURI)).
			With(zap.String("method", r.Method))

		//IP
		remoteAddr := r.RemoteAddr
		if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
			remoteAddr = realIP
		}
		l = l.With(zap.String("clientIP", remoteAddr))

		//X-Request-Id
		if reqID := r.Header.Get("X-Request-Id"); reqID != "" {
			l = l.With(zap.String("request_id", reqID))
		}

		h.ServeHTTP(w, r)

		l = l.With(zap.Duration("time", time.Now().Sub(start)))
		l.Info("")
	})
}
