package rest

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

func Recovery(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				zap.L().Error("Recovery", zap.Any("error", err))

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)

				errorResponse := &Error{
					Status:  http.StatusInternalServerError,
					Code:    ErrUnknown,
					Message: fmt.Sprint(err),
				}

				w.Write([]byte(errorResponse.Error()))
			}
		}()
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
