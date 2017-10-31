package restful

import (
	"github.com/NeuronFramework/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
	"net/http"
)

type errorResponder struct {
	status int
	err    error
}

func (e *errorResponder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(e.status)
	if err := producer.Produce(rw, e.err); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func Responder(err error) middleware.Responder {
	switch err.(type) {
	case *errors.Error:
		zap.L().Error("InternalServerError", zap.Error(err))
		return &errorResponder{status: err.(*errors.Error).Status, err: err}
	default:
		zap.L().Error("InternalServerError", zap.String("code", errors.ERROR_UNKNOWN), zap.Error(err))
		return &errorResponder{
			status: http.StatusInternalServerError,
			err:    &errors.Error{Status: http.StatusInternalServerError, Code: errors.ERROR_UNKNOWN, Message: err.Error()},
		}
	}
}
