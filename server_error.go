package restful

import (
	neuronError "github.com/NeuronFramework/errors"
	"github.com/go-openapi/errors"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func ServeError(rw http.ResponseWriter, r *http.Request, err error) {
	zap.L().Named("ServeError").Info("ServeError", zap.Error(err))
	rw.Header().Set("Content-Type", "application/json")
	switch e := err.(type) {
	case *neuronError.Error:
		rw.WriteHeader(e.Status)
		rw.Write([]byte(e.Error()))
		return
	case *errors.MethodNotAllowedError:
		rw.Header().Add("Allow", strings.Join(err.(*errors.MethodNotAllowedError).Allowed, ","))
		rw.WriteHeader(http.StatusMethodNotAllowed)
		if r == nil || r.Method != "HEAD" {
			apiError := neuronError.Error{
				Status:  http.StatusMethodNotAllowed,
				Code:    "MethodNotAllowed",
				Message: "Http method not allowed",
			}
			rw.Write([]byte(apiError.Error()))
		}
		return
	case errors.Error:
		if e == nil {
			rw.WriteHeader(http.StatusInternalServerError)
			apiError := neuronError.Unknown("e==nil")
			rw.Write([]byte(apiError.Error()))
			return
		}
		rw.WriteHeader(int(e.Code()))
		if r == nil || r.Method != "HEAD" {
			apiError := neuronError.Error{
				Status:  int(e.Code()),
				Code:    neuronError.ErrUnknown,
				Message: e.Error(),
			}
			rw.Write([]byte(apiError.Error()))
		}
		return
	default:
		rw.WriteHeader(http.StatusInternalServerError)
		if r == nil || r.Method != "HEAD" {
			apiError := neuronError.Unknown(e.Error())
			rw.Write([]byte(apiError.Error()))
		}
	}
}
