package utils

import (
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validate(data interface{}) error {
	return validate.Struct(data)
}