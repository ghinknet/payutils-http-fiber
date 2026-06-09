package httpfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func (i Instance) Get(path string, handler http.HandlerFunc) {
	i.Router.Get(path, adaptor.HTTPHandler(handler))
}

func (i Instance) Post(path string, handler http.HandlerFunc) {
	i.Router.Post(path, adaptor.HTTPHandler(handler))
}

func (i Instance) Put(path string, handler http.HandlerFunc) {
	i.Router.Put(path, adaptor.HTTPHandler(handler))
}

func (i Instance) Patch(path string, handler http.HandlerFunc) {
	i.Router.Patch(path, adaptor.HTTPHandler(handler))
}

func (i Instance) Delete(path string, handler http.HandlerFunc) {
	i.Router.Delete(path, adaptor.HTTPHandler(handler))
}

func (i Instance) Head(path string, handler http.HandlerFunc) {
	i.Router.Head(path, adaptor.HTTPHandler(handler))
}

func (i Instance) Options(path string, handler http.HandlerFunc) {
	i.Router.Options(path, adaptor.HTTPHandler(handler))
}

func (i Instance) Any(path string, handler http.HandlerFunc) {
	i.Router.All(path, adaptor.HTTPHandler(handler))
}
