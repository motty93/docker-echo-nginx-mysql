package validation

import "github.com/go-playground/validator/v10"

type Custom struct {
	Validator *validator.Validate
}

func (c *Custom) Validate(i interface{}) error {
	return c.Validator.Struct(i)
}
