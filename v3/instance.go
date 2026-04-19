package httpfiber

import (
	"github.com/ghinknet/payutils/v3/errors"
	"github.com/ghinknet/payutils/v3/model"
	"github.com/gofiber/fiber/v3"
)

type Instance struct {
	Router fiber.Router
}

type Driver struct{}

func (d Driver) NewInstance(instance any) (model.HttpInstance, error) {
	adaptedInstance := Instance{}

	// Assert type
	switch instance.(type) {
	case fiber.Router:
		adaptedInstance.Router = instance.(fiber.Router)
	default:
		return Instance{}, errors.ErrUnsupportedInstance
	}

	return adaptedInstance, nil
}
