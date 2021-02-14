package injector

import (
	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"
)

// CustomValidator CustomValidator
type CustomValidator struct {
	validator *validator.Validate
}

// NewValidator validatorの設定を行います.
func NewValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}

// Validate validate
func (c *CustomValidator) Validate(i interface{}) error {
	return c.validator.Struct(i)
}
