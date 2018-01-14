package restful

import (
	apiError "github.com/NeuronFramework/errors"
	"github.com/go-openapi/errors"
	"net/http"
	"strings"
)

func ServeError(rw http.ResponseWriter, r *http.Request, err error) {
	rw.Header().Set("Content-Type", "application/json")
	switch e := err.(type) {
	case *apiError.Error:
		rw.WriteHeader(e.Status)
		rw.Write([]byte(e.Error()))
		return
	case *errors.MethodNotAllowedError:
		rw.Header().Add("Allow", strings.Join(err.(*errors.MethodNotAllowedError).Allowed, ","))
		rw.WriteHeader(http.StatusMethodNotAllowed)
		if r == nil || r.Method != "HEAD" {
			apiError := apiError.Error{
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
			apiError := apiError.Unknown("e==nil")
			rw.Write([]byte(apiError.Error()))
			return
		}
		rw.WriteHeader(int(e.Code()))
		if r == nil || r.Method != "HEAD" {
			apiError := apiError.Error{
				Status:  int(e.Code()),
				Code:    apiError.ErrUnknown,
				Message: e.Error(),
			}
			rw.Write([]byte(apiError.Error()))
		}
		return
	default:
		rw.WriteHeader(http.StatusInternalServerError)
		if r == nil || r.Method != "HEAD" {
			apiError := apiError.Unknown(e.Error())
			rw.Write([]byte(apiError.Error()))
		}
	}
}
