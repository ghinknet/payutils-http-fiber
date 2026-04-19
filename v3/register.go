package httpfiber

import "github.com/ghinknet/payutils/v3/driver"

func init() {
	driver.RegisterHttp(Name, Driver{})
}
