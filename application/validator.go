package application

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() echo.Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
