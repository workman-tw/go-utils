package validator

import "github.com/go-playground/validator/v10"

// Validator - a validator for echo server
type Validator struct {
	validator *validator.Validate
}

// New - new validator
func New() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validate - To validate struct is right or not
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
