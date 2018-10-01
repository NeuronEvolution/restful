package rest

import (
	"go.uber.org/zap"
	"net/http"
)

func ServeError(rw http.ResponseWriter, r *http.Request, err error) {
	zap.L().Named("ServeError").Info("ServeError", zap.Error(err))

	rw.Header().Set("Content-Type", "application/json")

	e := Wrap(err)

	zap.L().Named("ServeError").Info("ServeErrorResponse", zap.Error(e))

	rw.WriteHeader(e.Status)
	rw.Write([]byte(e.Error()))

	return
}
