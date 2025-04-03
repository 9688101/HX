package valid

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

func GetValidate() *validator.Validate {
	return Validate
}

func ValidateStruct(v any) error {
	return Validate.Struct(v)
}

func ValidateVar(v any, tag string) error {
	return Validate.Var(v, tag)
}
