package httpfiber

import (
	"github.com/gofiber/fiber/v3"

	"go.gh.ink/payutils/v3/errors"
	"go.gh.ink/payutils/v3/model"
)

type Instance struct {
	Router fiber.Router
}

type Driver struct{}

func (d Driver) NewInstance(instance any) (model.HttpInstance, error) {
	router, ok := instance.(fiber.Router)
	if !ok {
		return Instance{}, errors.ErrUnsupportedInstance
	}
	return Instance{Router: router}, nil
}
