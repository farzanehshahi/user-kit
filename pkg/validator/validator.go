package validator

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() Validator {
	return Validator{validate: validator.New()}
}

// Validate validates input, based on the given model.
func (v Validator) Validate(model interface{}) (bool, error) {
	// todo: prepare error
	err := v.validate.Struct(model)
	if err != nil {
		return false, err
	}
	return true, nil
}
