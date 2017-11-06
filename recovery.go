package restful

import (
	"encoding/json"
	"fmt"
	"github.com/NeuronFramework/errors"
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

				errorResponse := &errors.Error{
					Status:  http.StatusInternalServerError,
					Code:    errors.ERROR_INTERNAL_EXCEPTION,
					Message: fmt.Sprint(err),
				}

				data, jsonError := json.Marshal(errorResponse)
				if jsonError != nil {
					zap.L().Error("Recovery", zap.Error(jsonError))
					w.Write([]byte(errorResponse.Error()))
					return
				}

				w.Write(data)
			}
		}()
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
