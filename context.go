package rest

import (
	"net/http"
	"time"
)

type Context struct {
	UserAgent string
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (ctx *Context) Done() <-chan struct{} {
	return nil
}

func (ctx *Context) Err() error {
	return nil
}

func (ctx *Context) Value(key interface{}) interface{} {
	return nil
}

func NewContext(r *http.Request) (ctx *Context) {
	ctx = &Context{}
	ctx.UserAgent = r.UserAgent()

	return ctx
}
