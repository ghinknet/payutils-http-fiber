package httpfiber

import (
	"net/http"

	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func (i Instance) Post(path string, handler http.HandlerFunc) {
	i.Router.Post(path, adaptor.HTTPHandler(handler))
}
