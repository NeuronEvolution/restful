package restful

import (
	"encoding/json"
	"github.com/NeuronFramework/errors"
	"net/http"
)

func ServeError(w http.ResponseWriter, r *http.Request, err error) {
	// todo: swagger errors
	e := &errors.Error{Status: http.StatusInternalServerError, Code: errors.ERROR_UNKNOWN, Message: err.Error()}
	s, _ := json.Marshal(e)
	w.WriteHeader(500)
	w.Write(s)
}
